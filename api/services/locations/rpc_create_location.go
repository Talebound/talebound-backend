package locations

import (
	"context"
	"database/sql"
	"github.com/the-medo/talebound-backend/api/servicecore"
	"github.com/the-medo/talebound-backend/converters"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *ServiceLocations) CreateLocation(ctx context.Context, request *pb.CreateLocationRequest) (*pb.Location, error) {
	violations := validateCreateLocation(request)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	_, err := server.CheckModuleIdPermissions(ctx, request.GetModuleId(), &servicecore.ModulePermission{
		NeedsEntityPermission: &[]db.EntityType{db.EntityTypeLocation},
	})

	location, err := server.Store.CreateLocation(ctx, db.CreateLocationParams{
		Name: request.GetName(),
		Description: sql.NullString{
			String: request.GetDescription(),
			Valid:  true,
		},
		ThumbnailImageID: sql.NullInt32{
			Int32: request.GetThumbnailImageId(),
			Valid: request.ThumbnailImageId != nil,
		},
	})

	if err != nil {
		return nil, err
	}

	_, err = server.Store.CreateEntity(ctx, db.CreateEntityParams{
		Type:     db.EntityTypeLocation,
		ModuleID: request.GetModuleId(),
		LocationID: sql.NullInt32{
			Int32: location.ID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	rsp := converters.ConvertLocation(location)

	return rsp, nil
}

func validateCreateLocation(req *pb.CreateLocationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUniversalId(req.GetModuleId()); err != nil {
		violations = append(violations, e.FieldViolation("module_id", err))
	}

	if err := validator.ValidateLocationName(req.GetName()); err != nil {
		violations = append(violations, e.FieldViolation("name", err))
	}

	if req.Description != nil {
		if err := validator.ValidateLocationDescription(req.GetDescription()); err != nil {
			violations = append(violations, e.FieldViolation("description", err))
		}
	}

	if req.ThumbnailImageId != nil {
		if err := validator.ValidateImageId(req.GetThumbnailImageId()); err != nil {
			violations = append(violations, e.FieldViolation("thumbnail_image_id", err))
		}
	}

	return violations
}
