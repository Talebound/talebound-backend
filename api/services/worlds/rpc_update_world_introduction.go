package worlds

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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServiceWorlds) UpdateWorldIntroduction(ctx context.Context, req *pb.UpdateWorldIntroductionRequest) (*pb.ViewPost, error) {
	violations := validateUpdateWorldIntroductionRequest(req)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	authPayload, module, err := server.CheckModuleTypePermissions(ctx, db.ModuleTypeWorld, req.GetWorldId(), &servicecore.ModulePermission{
		NeedsEntityPermission: &[]db.EntityType{db.EntityTypePost},
	})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "failed update introduction: %v", err)
	}

	world, err := server.Store.GetWorldByID(ctx, req.GetWorldId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to save changes - world not found: %s", err)
	}

	if !world.DescriptionPostID.Valid {
		//create new post
		createPostArg := db.CreatePostParams{
			UserID:    authPayload.UserId,
			Title:     "World introduction",
			Content:   req.GetContent(),
			IsDraft:   false,
			IsPrivate: false,
		}

		post, err := server.Store.CreatePost(ctx, createPostArg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create post: %s", err)
		}

		//if it is post (so first time introduction is created), we put it into users table
		updateWorldArg := db.UpdateWorldParams{
			WorldID: req.WorldId,
			DescriptionPostID: sql.NullInt32{
				Int32: post.ID,
				Valid: true,
			},
		}
		_, err = server.Store.UpdateWorld(ctx, updateWorldArg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update world: %s", err)
		}

		_, err = server.Store.CreateEntity(ctx, db.CreateEntityParams{
			Type:     db.EntityTypePost,
			ModuleID: module.ID,
			PostID: sql.NullInt32{
				Int32: post.ID,
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create entity post: %s", err)
		}

		viewPost, err := server.Store.GetPostById(ctx, post.ID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get post: %s", err)
		}

		return converters.ConvertViewPost(viewPost), nil
	} else {
		//update existing post
		arg := db.UpdatePostParams{
			PostID: world.DescriptionPostID.Int32,
			LastUpdatedUserID: sql.NullInt32{
				Int32: authPayload.UserId,
				Valid: true,
			},
			Content: sql.NullString{
				String: req.GetContent(),
				Valid:  true,
			},
		}
		post, err := server.Store.UpdatePost(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update post: %s", err)
		}

		viewPost, err := server.Store.GetPostById(ctx, post.ID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get post: %s", err)
		}

		return converters.ConvertViewPost(viewPost), nil
	}
}

func validateUpdateWorldIntroductionRequest(req *pb.UpdateWorldIntroductionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateWorldId(req.GetWorldId()); err != nil {
		violations = append(violations, e.FieldViolation("world_id", err))
	}

	if err := validator.ValidatePostContent(req.GetContent()); err != nil {
		violations = append(violations, e.FieldViolation("content", err))
	}

	return violations
}