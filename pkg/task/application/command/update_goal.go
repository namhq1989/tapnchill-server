package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type UpdateGoalHandler struct {
	goalRepository domain.GoalRepository
	service        domain.Service
}

func NewUpdateGoalHandler(goalRepository domain.GoalRepository, service domain.Service) UpdateGoalHandler {
	return UpdateGoalHandler{
		goalRepository: goalRepository,
		service:        service,
	}
}

func (h UpdateGoalHandler) UpdateGoal(ctx *appcontext.AppContext, performerID, goalID string, req dto.UpdateGoalRequest) (*dto.UpdateGoalResponse, error) {
	ctx.Logger().Info("new update goal request", appcontext.Fields{
		"performerID": performerID, "goalID": goalID,
		"name": req.Name, "description": req.Description,
	})

	goal, err := h.service.GetGoalByID(ctx, goalID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("update goal")
	if err = goal.SetName(req.Name); err != nil {
		ctx.Logger().Error("failed to update goal name", err, appcontext.Fields{})
		return nil, err
	}
	goal.SetDescription(req.Description)

	ctx.Logger().Text("update goal in db")
	if err = h.goalRepository.Update(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to update goal in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update goal request")
	return &dto.UpdateGoalResponse{
		ID: goal.ID,
	}, nil
}
