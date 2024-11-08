package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type UpdateTaskHandler struct {
	taskRepository domain.TaskRepository
}

func NewUpdateTaskHandler(taskRepository domain.TaskRepository) UpdateTaskHandler {
	return UpdateTaskHandler{
		taskRepository: taskRepository,
	}
}

func (h UpdateTaskHandler) UpdateTask(ctx *appcontext.AppContext, performerID, taskID string, req dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {
	ctx.Logger().Info("new update task request", appcontext.Fields{
		"performerID": performerID, "taskID": taskID,
		"name": req.Name, "description": req.Description, "dueDate": req.DueDate,
	})

	ctx.Logger().Text("find task in db")
	task, err := h.taskRepository.FindByID(ctx, taskID)
	if err != nil {
		ctx.Logger().Error("failed to find task in db", err, appcontext.Fields{})
		return nil, err
	}
	if task == nil {
		ctx.Logger().ErrorText("task not found, respond")
		return nil, apperrors.Common.NotFound
	}

	if task.UserID != performerID {
		ctx.Logger().ErrorText("task author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("update task")
	if err = task.SetName(req.Name); err != nil {
		ctx.Logger().Error("failed to update task name", err, appcontext.Fields{})
		return nil, err
	}
	task.SetDescription(req.Description)
	task.SetDueDate(req.DueDate)

	ctx.Logger().Text("update task in db")
	if err = h.taskRepository.Update(ctx, *task); err != nil {
		ctx.Logger().Error("failed to update task in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update task request")
	return &dto.UpdateTaskResponse{
		ID: task.ID,
	}, nil
}
