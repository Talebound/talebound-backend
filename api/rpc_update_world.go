package api

import (
	"context"
	"database/sql"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateWorld(ctx context.Context, req *pb.UpdateWorldRequest) (*pb.World, error) {
	violations := validateUpdateWorldRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.CheckWorldAdmin(ctx, req.GetWorldId(), true)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "failed to update world: %v", err)
	}

	arg := db.UpdateWorldParams{
		WorldID: req.GetWorldId(),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		BasedOn: sql.NullString{
			String: req.GetBasedOn(),
			Valid:  req.BasedOn != nil,
		},
		ShortDescription: sql.NullString{
			String: req.GetShortDescription(),
			Valid:  req.ShortDescription != nil,
		},
		Public: sql.NullBool{
			Bool:  req.GetPublic(),
			Valid: req.Public != nil,
		},
		DescriptionPostID: sql.NullInt32{
			Int32: req.GetDescriptionPostId(),
			Valid: req.DescriptionPostId != nil,
		},
	}

	changesMade := arg.Name.Valid || arg.BasedOn.Valid || arg.ShortDescription.Valid || arg.Public.Valid || arg.DescriptionPostID.Valid
	if changesMade {
		_, err = server.store.UpdateWorld(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update world: %v", err)
		}
	}

	argImages := db.UpdateWorldImagesParams{
		WorldID: req.GetWorldId(),
		ThumbnailImgID: sql.NullInt32{
			Int32: req.GetImageThumbnailId(),
			Valid: req.ImageThumbnailId != nil,
		},
		HeaderImgID: sql.NullInt32{
			Int32: req.GetImageHeaderId(),
			Valid: req.ImageHeaderId != nil,
		},
		AvatarImgID: sql.NullInt32{
			Int32: req.GetImageAvatarId(),
			Valid: req.ImageAvatarId != nil,
		},
	}

	changesMade = argImages.ThumbnailImgID.Valid || argImages.HeaderImgID.Valid || argImages.AvatarImgID.Valid

	if changesMade {
		_, err = server.store.UpdateWorldImages(ctx, argImages)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update world: %v", err)
		}
	}

	world, err := server.store.GetWorldByID(ctx, req.GetWorldId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve updated world: %v", err)
	}

	rsp := convertWorld(world)

	return rsp, nil
}

func validateUpdateWorldRequest(req *pb.UpdateWorldRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := validator.ValidateWorldId(req.GetWorldId()); err != nil {
		violations = append(violations, FieldViolation("world_id", err))
	}

	if req.Name != nil {
		if err := validator.ValidateWorldName(req.GetName()); err != nil {
			violations = append(violations, FieldViolation("name", err))
		}
	}

	if req.ShortDescription != nil {
		if err := validator.ValidateWorldShortDescription(req.GetShortDescription()); err != nil {
			violations = append(violations, FieldViolation("short_description", err))
		}
	}

	if req.BasedOn != nil {
		if err := validator.ValidateWorldBasedOn(req.GetBasedOn()); err != nil {
			violations = append(violations, FieldViolation("based_on", err))
		}
	}

	if req.DescriptionPostId != nil {
		if err := validator.ValidatePostId(req.GetDescriptionPostId()); err != nil {
			violations = append(violations, FieldViolation("description_post_id", err))
		}
	}

	if req.ImageAvatarId != nil {
		if err := validator.ValidateImageId(req.GetImageAvatarId()); err != nil {
			violations = append(violations, FieldViolation("image_avatar_id", err))
		}
	}

	if req.ImageThumbnailId != nil {
		if err := validator.ValidateImageId(req.GetImageThumbnailId()); err != nil {
			violations = append(violations, FieldViolation("image_thumbnail_id", err))
		}
	}

	if req.ImageHeaderId != nil {
		if err := validator.ValidateImageId(req.GetImageHeaderId()); err != nil {
			violations = append(violations, FieldViolation("image_header_id", err))
		}
	}

	return violations
}