package api

import (
	"context"
	"database/sql"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/util"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	viewUser, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "viewUser not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find viewUser: %v", err)
	}

	err = util.CheckPassword(req.Password, viewUser.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password: %v", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		viewUser.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %v", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		viewUser.ID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %v", err)
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       viewUser.ID,
		Username:     viewUser.Username,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %v", err)
	}

	user := convertViewUserToUser(viewUser)
	rsp := &pb.LoginUserResponse{
		User:                  convertUserGetImage(server, ctx, user),
		SessionId:             session.ID.String(),
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	md := metadata.Pairs(
		"X-Access-Token", accessToken,
		"X-Access-Token-Expires-At", accessPayload.ExpiredAt.Format(util.TimeLayout),
		"X-Refresh-Token", refreshToken,
		"X-Refresh-Token-Expires-At", refreshPayload.ExpiredAt.Format(util.TimeLayout),
	)
	err = grpc.SendHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, FieldViolation("username", err))
	}

	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, FieldViolation("password", err))
	}

	return violations
}
