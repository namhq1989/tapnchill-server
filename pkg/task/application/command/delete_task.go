package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type DeleteTaskHandler struct {
	taskRepository domain.TaskRepository
	goalRepository domain.GoalRepository
	service        domain.Service
}

func NewDeleteTaskHandler(taskRepository domain.TaskRepository, goalRepository domain.GoalRepository, service domain.Service) DeleteTaskHandler {
	return DeleteTaskHandler{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
		service:        service,
	}
}

func (h DeleteTaskHandler) DeleteTask(ctx *appcontext.AppContext, performerID, taskID string, _ dto.DeleteTaskRequest) (*dto.DeleteTaskResponse, error) {
	ctx.Logger().Info("new delete task request", appcontext.Fields{"performerID": performerID, "taskID": taskID})

	task, err := h.service.GetTaskByID(ctx, taskID, performerID)
	if err != nil {
		return nil, err
	}

	goal, err := h.service.GetGoalByID(ctx, task.GoalID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("delete task in db")
	if err = h.taskRepository.Delete(ctx, taskID); err != nil {
		ctx.Logger().Error("failed to delete task in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update goal stats")
	goal.AdjustTotalTask(-1)
	if task.IsCompleted {
		goal.AdjustTotalCompletedTask(-1)
	}

	ctx.Logger().Text("update goal in db")
	if err = h.goalRepository.Update(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to update goal in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done delete task request")
	return &dto.DeleteTaskResponse{}, nil
}
