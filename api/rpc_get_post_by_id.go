package api

import (
	"context"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.Post, error) {
	violations := validateGetPostByIdRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	post, err := server.store.GetPostById(ctx, req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get post: %v", err)
	}

	rsp := convertViewPost(post)

	return rsp, nil
}

func validateGetPostByIdRequest(req *pb.GetPostByIdRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidatePostId(req.GetPostId()); err != nil {
		violations = append(violations, FieldViolation("post_id", err))
	}

	return violations
}