package api

import (
	"context"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserRoles(ctx context.Context, req *pb.GetUserRolesRequest) (*pb.GetUserRolesResponse, error) {

	violations := validateGetUserRoles(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	roles, err := server.store.GetUserRoles(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user roles: %v", err)
	}

	rsp := &pb.GetUserRolesResponse{
		Role: make([]*pb.Role, len(roles)),
	}

	for i, role := range roles {
		rsp.Role[i] = &pb.Role{
			Id:          role.RoleID,
			Name:        role.RoleName,
			Description: role.RoleDescription,
		}
	}

	return rsp, nil
}

func validateGetUserRoles(req *pb.GetUserRolesRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := validator.ValidateUserId(req.GetUserId()); err != nil {
		violations = append(violations, FieldViolation("user_id", err))
	}

	return violations
}