package api

import (
	"context"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetUserImages(ctx context.Context, req *pb.GetUserImagesRequest) (*pb.GetImagesResponse, error) {
	violations := validateGetUserImagesRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	limit := req.GetLimit()
	offset := req.GetOffset()

	getImagesReq := &pb.GetImagesRequest{
		UserId:      &req.UserId,
		ImageTypeId: req.ImageTypeId,
		Limit:       &limit,
		Offset:      &offset,
	}

	return server.GetImages(ctx, getImagesReq)
}

func validateGetUserImagesRequest(req *pb.GetUserImagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := validator.ValidateUserId(req.GetUserId()); err != nil {
		violations = append(violations, FieldViolation("user_id", err))
	}

	if req.ImageTypeId != nil {
		if err := validator.ValidateImageTypeId(req.GetImageTypeId()); err != nil {
			violations = append(violations, FieldViolation("post_type_id", err))
		}
	}

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