package api

import (
	"context"
	"database/sql"
	"github.com/the-medo/talebound-backend/consts"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUserIntroduction(ctx context.Context, req *pb.UpdateUserIntroductionRequest) (*pb.Post, error) {
	violations := validateUpdateUserIntroductionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.authorizeUserCookie(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if req.GetUserId() != authPayload.UserId {
		err := server.CheckUserRole(ctx, []pb.RoleType{pb.RoleType_admin})
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "unable to save changes - you are not creator or admin: %v", err)
		}
	}

	user, err := server.store.GetUserById(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to save changes - user not found: %s", err)
	}

	postType, err := server.store.GetPostTypeById(ctx, consts.PostTypeUserIntroduction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get post type: %v", err)
	}

	if !postType.Draftable && req.GetSaveAsDraft() {
		return nil, status.Errorf(codes.InvalidArgument, "cannot save this post as draft")
	}

	if !user.IntroductionPostID.Valid {
		//create new post
		createPostArg := db.CreatePostParams{
			UserID:     req.GetUserId(),
			Title:      "User introduction",
			PostTypeID: consts.PostTypeUserIntroduction,
			Content:    req.GetContent(),
			IsDraft:    req.GetSaveAsDraft(),
			IsPrivate:  false,
		}

		post, err := server.store.CreatePost(ctx, createPostArg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create post: %s", err)
		}

		//if it is post (so first time introduction is created), we put it into users table
		updateUserArg := db.UpdateUserParams{
			ID: req.UserId,
			IntroductionPostID: sql.NullInt32{
				Int32: post.ID,
				Valid: true,
			},
		}
		_, err = server.store.UpdateUser(ctx, updateUserArg)

		return convertPostAndPostType(post, postType), nil
	} else {
		//update existing post
		arg := db.UpdatePostParams{
			PostID: user.IntroductionPostID.Int32,
			LastUpdatedUserID: sql.NullInt32{
				Int32: authPayload.UserId,
				Valid: true,
			},
			Content: sql.NullString{
				String: req.GetContent(),
				Valid:  true,
			},
			IsDraft: sql.NullBool{
				Bool:  req.GetSaveAsDraft(),
				Valid: req.SaveAsDraft != nil,
			},
		}
		post, err := server.store.UpdatePost(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update post: %s", err)
		}

		return convertPostAndPostType(post, postType), nil
	}
}

func validateUpdateUserIntroductionRequest(req *pb.UpdateUserIntroductionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUserId(req.GetUserId()); err != nil {
		violations = append(violations, FieldViolation("user_id", err))
	}

	if err := validator.ValidatePostContent(req.GetContent()); err != nil {
		violations = append(violations, FieldViolation("content", err))
	}

	return violations
}