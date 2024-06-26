package maps

import (
	"context"
	"fmt"
	"github.com/the-medo/talebound-backend/converters"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *ServiceMaps) CreateMapLayer(ctx context.Context, request *pb.CreateMapLayerRequest) (*pb.ViewMapLayer, error) {
	violations := validateCreateMapLayer(request)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	_, err := server.CheckEntityTypePermissions(ctx, db.EntityTypeMap, request.GetMapId(), nil)
	if err != nil {
		return nil, err
	}

	maxLayerPositionForMap, err := server.Store.GetMaxMapLayerPosition(ctx, request.GetMapId())
	if err != nil {
		return nil, err
	}

	argLayer := db.CreateMapLayerParams{
		MapID:   request.GetMapId(),
		Name:    request.GetName(),
		ImageID: request.GetImageId(),
		Enabled: request.GetEnabled(),
	}

	if request.Position != nil {
		if request.GetPosition() > maxLayerPositionForMap+1 {
			return nil, fmt.Errorf("position out of range")
		}
		argLayer.Position = request.GetPosition()

		err := server.Store.IncreaseMapLayerPositions(ctx, db.IncreaseMapLayerPositionsParams{
			MapID:    request.GetMapId(),
			Position: request.GetPosition(),
		})
		if err != nil {
			return nil, err
		}
	} else {
		argLayer.Position = maxLayerPositionForMap + 1
	}

	mapRow, err := server.Store.GetMapById(ctx, request.GetMapId())
	if err != nil {
		return nil, err
	}
	imageRow, err := server.Store.GetImageById(ctx, request.GetImageId())
	if err != nil {
		return nil, err
	}

	if (mapRow.Width != imageRow.Width) || (mapRow.Height != imageRow.Height) {
		return nil, e.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{
			{
				Field:       "image_id",
				Description: "size of map layer image must be the same size as the map",
			},
		})
	}

	newLayer, err := server.Store.CreateMapLayer(ctx, argLayer)
	if err != nil {
		return nil, err
	}

	viewMapLayer, err := server.Store.GetMapLayerByID(ctx, newLayer.ID)
	if err != nil {
		return nil, err
	}

	rsp := converters.ConvertViewMapLayer(viewMapLayer)

	return rsp, nil
}

func validateCreateMapLayer(req *pb.CreateMapLayerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateMapId(req.GetMapId()); err != nil {
		violations = append(violations, e.FieldViolation("map_id", err))
	}

	if err := validator.ValidateUniversalName(req.GetName()); err != nil {
		violations = append(violations, e.FieldViolation("name", err))
	}

	if err := validator.ValidateImageId(req.GetImageId()); err != nil {
		violations = append(violations, e.FieldViolation("image_id", err))
	}

	return violations
}
