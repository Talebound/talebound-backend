package maps

import (
	"context"
	"github.com/the-medo/talebound-backend/api/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *ServiceMaps) DeleteMapPinType(ctx context.Context, request *pb.DeleteMapPinTypeRequest) (*emptypb.Empty, error) {
	violations := validateDeleteMapPinType(request)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	err := server.CheckMapAccess(ctx, request.GetMapId(), false)
	if err != nil {
		return nil, err
	}

	err = server.Store.DeleteMapPinType(ctx, request.GetPinTypeId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func validateDeleteMapPinType(req *pb.DeleteMapPinTypeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateMapId(req.GetMapId()); err != nil {
		violations = append(violations, e.FieldViolation("map_id", err))
	}

	if err := validator.ValidateUniversalId(req.GetPinTypeId()); err != nil {
		violations = append(violations, e.FieldViolation("pin_type_id", err))
	}

	return violations
}