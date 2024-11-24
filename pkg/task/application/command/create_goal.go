package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type CreateGoalHandler struct {
	goalRepository domain.GoalRepository
	userHub        domain.UserHub
}

func NewCreateGoalHandler(goalRepository domain.GoalRepository, userHub domain.UserHub) CreateGoalHandler {
	return CreateGoalHandler{
		goalRepository: goalRepository,
		userHub:        userHub,
	}
}

func (h CreateGoalHandler) CreateGoal(ctx *appcontext.AppContext, performerID string, req dto.CreateGoalRequest) (*dto.CreateGoalResponse, error) {
	ctx.Logger().Info("new create goal request", appcontext.Fields{"performerID": performerID, "name": req.Name, "description": req.Description})

	ctx.Logger().Text("get user goal quota")
	quota, err := h.userHub.GetGoalQuota(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get user goal quota", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count user total goals")
	totalGoals, err := h.goalRepository.CountByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to count user total goals", err, appcontext.Fields{})
		return nil, err
	}

	if totalGoals >= quota {
		ctx.Logger().Error("user goal quota exceeded", err, appcontext.Fields{"quota": quota, "total": totalGoals})
		return nil, apperrors.User.ResourceLimitReached
	}

	ctx.Logger().Text("create new goal model")
	goal, err := domain.NewGoal(performerID, req.Name, req.Description)
	if err != nil {
		ctx.Logger().Error("failed to create new goal model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist goal in db")
	if err = h.goalRepository.Create(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to persist goal in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create goal request")
	return &dto.CreateGoalResponse{
		ID: goal.ID,
	}, nil
}
