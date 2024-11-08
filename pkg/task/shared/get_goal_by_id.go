package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
)

func (s Service) GetGoalByID(ctx *appcontext.AppContext, goalID, userID string) (*domain.Goal, error) {
	ctx.Logger().Text("find goal in db")
	goal, err := s.goalRepository.FindByID(ctx, goalID)
	if err != nil {
		ctx.Logger().Error("failed to find goal in db", err, appcontext.Fields{})
		return nil, err
	}
	if goal == nil {
		ctx.Logger().ErrorText("goal not found, respond")
		return nil, apperrors.Common.NotFound
	}
	if goal.UserID != userID {
		ctx.Logger().ErrorText("goal author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	return goal, nil
}
