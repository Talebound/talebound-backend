package api

import (
	"context"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	violations := validateGetUsers(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	limit, offset := GetDefaultQueryBoundaries(req.GetLimit(), req.GetOffset())

	users, err := server.store.GetUsers(ctx, db.GetUsersParams{
		PageLimit:  limit,
		PageOffset: offset,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	rsp := &pb.GetUsersResponse{
		Users: make([]*pb.User, len(users)),
	}

	for i, user := range users {
		rsp.Users[i] = convertUserRowWithImage(user)
	}

	return rsp, nil
}

func validateGetUsers(req *pb.GetUsersRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Limit != nil {
		if err := validator.ValidateLimit(req.GetLimit()); err != nil {
			violations = append(violations, FieldViolation("limit", err))
		}
	}

	if req.Offset != nil {
		if err := validator.ValidateOffset(req.GetOffset()); err != nil {
			violations = append(violations, FieldViolation("offset", err))
		}
	}

	return violations
}