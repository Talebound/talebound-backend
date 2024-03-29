package evaluations

import (
	"context"
	db "github.com/the-medo/talebound-backend/db/sqlc"
	"github.com/the-medo/talebound-backend/e"
	"github.com/the-medo/talebound-backend/pb"
	"github.com/the-medo/talebound-backend/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

func (server *ServiceEvaluations) GetAverageUserEvaluationsByType(ctx context.Context, req *pb.GetAverageUserEvaluationsByTypeRequest) (*pb.GetAverageUserEvaluationsByTypeResponse, error) {
	violations := validateGetAverageUserEvaluationsByType(req)
	if violations != nil {
		return nil, e.InvalidArgumentError(violations)
	}

	evaluationType := db.EvaluationType(req.GetType())

	arg := db.GetAverageUserEvaluationsByTypeParams{
		UserID:         req.GetUserId(),
		EvaluationType: evaluationType,
	}

	avgEvaluationVotes, err := server.Store.GetAverageUserEvaluationsByType(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get evaluation averages: %v", err)
	}

	rsp := &pb.GetAverageUserEvaluationsByTypeResponse{
		AverageEvaluationVote: make([]*pb.AverageEvaluationVote, len(avgEvaluationVotes)),
	}

	for i, eval := range avgEvaluationVotes {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to convert evaluation type: %v", err)
		}

		rsp.AverageEvaluationVote[i] = &pb.AverageEvaluationVote{
			EvaluationId: eval.EvaluationID,
			UserId:       req.GetUserId(),
			Name:         eval.Name,
			Description:  eval.Description,
			Type:         string(eval.EvaluationType),
			Average:      math.Round(eval.AvgValue*100) / 100,
		}
	}

	return rsp, nil
}

func validateGetAverageUserEvaluationsByType(req *pb.GetAverageUserEvaluationsByTypeRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := validator.ValidateUserId(req.GetUserId()); err != nil {
		violations = append(violations, e.FieldViolation("user_id", err))
	}

	if err := validator.ValidateEvaluationType(req.GetType()); err != nil {
		violations = append(violations, e.FieldViolation("type", err))
	}

	return violations
}
