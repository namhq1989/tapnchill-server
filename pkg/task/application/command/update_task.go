package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type UpdateTaskHandler struct {
	taskRepository domain.TaskRepository
	service        domain.Service
}

func NewUpdateTaskHandler(taskRepository domain.TaskRepository, service domain.Service) UpdateTaskHandler {
	return UpdateTaskHandler{
		taskRepository: taskRepository,
		service:        service,
	}
}

func (h UpdateTaskHandler) UpdateTask(ctx *appcontext.AppContext, performerID, taskID string, req dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {
	ctx.Logger().Info("new update task request", appcontext.Fields{
		"performerID": performerID, "taskID": taskID,
		"name": req.Name, "description": req.Description, "dueDate": req.DueDate,
	})

	task, err := h.service.GetTaskByID(ctx, taskID, performerID)
	if err != nil {
		return nil, err
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
