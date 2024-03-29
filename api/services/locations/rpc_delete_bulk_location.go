package locations

import (
	"context"
	"github.com/the-medo/talebound-backend/api/servicecore"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *ServiceLocations) DeleteBulkLocation(ctx context.Context, request *pb.DeleteBulkLocationRequest) (*emptypb.Empty, error) {
	violations := validateDeleteBulkLocation(request)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	permissionsNeeded := servicecore.ModulePermission{
		NeedsEntityPermission: &[]db.EntityType{db.EntityTypeLocation},
	}
	//TODO: probably not the best way to do this
	for _, locationId := range request.GetLocationIds() {

		_, err := server.CheckEntityTypePermissions(ctx, db.EntityTypeLocation, locationId, &permissionsNeeded)
		if err != nil {
			return nil, err
		}

		err = server.Store.DeleteLocation(ctx, locationId)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func validateDeleteBulkLocation(req *pb.DeleteBulkLocationRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	for _, locationId := range req.GetLocationIds() {
		if err := validator.ValidateLocationId(locationId); err != nil {
			violations = append(violations, e.FieldViolation("location_id", err))
		}
	}
	return violations
}
