package api

import (
	"context"
	"github.com/the-medo/talebound-backend/api/converters"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAvailableWorldTag(ctx context.Context, req *pb.CreateAvailableWorldTagRequest) (*pb.Tag, error) {

	violations := validateCreateAvailableWorldTagRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err := server.CheckUserRole(ctx, []pb.RoleType{pb.RoleType_admin})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "can not create new tag - you are not admin: %v", err)
	}

	tag, err := server.store.CreateWorldTagAvailable(ctx, req.GetTag())
	if err != nil {
		return nil, err
	}

	rsp := converters.ConvertTag(tag)

	return rsp, nil
}

func validateCreateAvailableWorldTagRequest(req *pb.CreateAvailableWorldTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := validator.ValidateTag(req.GetTag()); err != nil {
		violations = append(violations, FieldViolation("tag", err))
	}

	return violations
}