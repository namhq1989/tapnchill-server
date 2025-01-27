package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type DeleteGoalHandler struct {
	goalRepository domain.GoalRepository
	service        domain.Service
}

func NewDeleteGoalHandler(goalRepository domain.GoalRepository, service domain.Service) DeleteGoalHandler {
	return DeleteGoalHandler{
		goalRepository: goalRepository,
		service:        service,
	}
}

func (h DeleteGoalHandler) DeleteGoal(ctx *appcontext.AppContext, performerID, goalID string, _ dto.DeleteGoalRequest) (*dto.DeleteGoalResponse, error) {
	ctx.Logger().Info("new delete goal request", appcontext.Fields{"performerID": performerID, "goalID": goalID})

	goal, err := h.service.GetGoalByID(ctx, goalID, performerID)
	if err != nil {
		return nil, err
	}
	if goal.Stats.TotalTask > 0 {
		ctx.Logger().ErrorText("goal has tasks, respond")
		return nil, apperrors.Task.GoalDeleteErrorTasksRemaining
	}

	ctx.Logger().Text("delete goal in db")
	if err = h.goalRepository.Delete(ctx, goalID); err != nil {
		ctx.Logger().Error("failed to delete goal in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done delete goal request")
	return &dto.DeleteGoalResponse{}, nil
}
