package api

import (
	"context"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/util"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetEvaluationById(ctx context.Context, req *pb.GetEvaluationByIdRequest) (*pb.GetEvaluationByIdResponse, error) {

	violations := validateGetEvaluationById(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	evaluation, err := server.store.GetEvaluationById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user roles: %v", err)
	}

	evaluationType, err := util.StringToEvaluationType(string(evaluation.EvaluationType))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert evaluation type: %v", err)
	}

	rsp := &pb.GetEvaluationByIdResponse{
		Evaluation: &pb.Evaluation{
			Id:          evaluation.ID,
			Name:        evaluation.Name,
			Description: evaluation.Description,
			Type:        evaluationType,
		},
	}
	return rsp, nil
}

func validateGetEvaluationById(req *pb.GetEvaluationByIdRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEvaluationId(req.GetId()); err != nil {
		violations = append(violations, FieldViolation("id", err))
	}
	return violations
}