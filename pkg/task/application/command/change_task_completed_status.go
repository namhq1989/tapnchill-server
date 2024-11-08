package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type ChangeTaskCompletedStatusHandler struct {
	taskRepository domain.TaskRepository
	goalRepository domain.GoalRepository
	service        domain.Service
}

func NewChangeTaskCompletedStatusHandler(
	taskRepository domain.TaskRepository,
	goalRepository domain.GoalRepository,
	service domain.Service,
) ChangeTaskCompletedStatusHandler {
	return ChangeTaskCompletedStatusHandler{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
		service:        service,
	}
}

func (h ChangeTaskCompletedStatusHandler) ChangeTaskCompletedStatus(ctx *appcontext.AppContext, performerID, taskID string, req dto.ChangeTaskCompletedStatusRequest) (*dto.ChangeTaskCompletedStatusResponse, error) {
	ctx.Logger().Info("new change task completed status request", appcontext.Fields{
		"performerID": performerID, "taskID": taskID,
		"completed": req.Completed,
	})

	task, err := h.service.GetTaskByID(ctx, taskID, performerID)
	if err != nil {
		return nil, err
	}

	if task.IsCompleted == req.Completed {
		ctx.Logger().Text("task completed status not changed, respond")
		return &dto.ChangeTaskCompletedStatusResponse{
			Completed: task.IsCompleted,
		}, nil
	}

	goal, err := h.service.GetGoalByID(ctx, task.GoalID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("change task completed status")
	task.SetCompleted(req.Completed)

	ctx.Logger().Text("update task in db")
	if err = h.taskRepository.Update(ctx, *task); err != nil {
		ctx.Logger().Error("failed to update task in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("adjust goal stats")
	adjustValue := 1
	if !req.Completed {
		adjustValue = -1
	}
	goal.AdjustTotalCompletedTask(adjustValue)

	ctx.Logger().Text("update goal in db")
	if err = h.goalRepository.Update(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to update goal in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done change task completed status request")
	return &dto.ChangeTaskCompletedStatusResponse{
		Completed: task.IsCompleted,
	}, nil
}
