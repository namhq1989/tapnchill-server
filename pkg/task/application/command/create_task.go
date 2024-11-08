package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type CreateTaskHandler struct {
	taskRepository domain.TaskRepository
	goalRepository domain.GoalRepository
}

func NewCreateTaskHandler(taskRepository domain.TaskRepository, goalRepository domain.GoalRepository) CreateTaskHandler {
	return CreateTaskHandler{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
	}
}

func (h CreateTaskHandler) CreateTask(ctx *appcontext.AppContext, performerID string, req dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	ctx.Logger().Info("new create task request", appcontext.Fields{
		"performerID": performerID, "goalID": req.GoalID,
		"name": req.Name, "description": req.Description, "dueDate": req.DueDate,
	})

	ctx.Logger().Text("find goal in db")
	goal, err := h.goalRepository.FindByID(ctx, req.GoalID)
	if err != nil {
		ctx.Logger().Error("failed to find goal in db", err, appcontext.Fields{})
		return nil, err
	}
	if goal == nil {
		ctx.Logger().ErrorText("goal not found, respond")
		return nil, apperrors.Common.NotFound
	}
	if goal.UserID != performerID {
		ctx.Logger().ErrorText("goal author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("create new task model")
	task, err := domain.NewTask(performerID, req.GoalID, req.Name, req.Description, req.DueDate)
	if err != nil {
		ctx.Logger().Error("failed to create new task model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist task in db")
	if err = h.taskRepository.Create(ctx, *task); err != nil {
		ctx.Logger().Error("failed to persist task in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("increase goal stats")
	goal.AdjustTotalTask(1)
	ctx.Logger().Text("update goal in db")
	if err = h.goalRepository.Update(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to update goal in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done create task request")
	return &dto.CreateTaskResponse{
		ID: task.ID,
	}, nil
}
