package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type ChangeTaskStatusHandler struct {
	taskRepository domain.TaskRepository
	goalRepository domain.GoalRepository
	service        domain.Service
}

func NewChangeTaskStatusHandler(
	taskRepository domain.TaskRepository,
	goalRepository domain.GoalRepository,
	service domain.Service,
) ChangeTaskStatusHandler {
	return ChangeTaskStatusHandler{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
		service:        service,
	}
}

func (h ChangeTaskStatusHandler) ChangeTaskStatus(ctx *appcontext.AppContext, performerID, taskID string, req dto.ChangeTaskStatusRequest) (*dto.ChangeTaskStatusResponse, error) {
	ctx.Logger().Info("new change task status request", appcontext.Fields{
		"performerID": performerID, "taskID": taskID, "status": req.Status,
	})

	ctx.Logger().Text("check status")
	status := domain.ToTaskStatus(req.Status)
	if !status.IsValid() {
		ctx.Logger().ErrorText("invalid status, respond")
		return nil, apperrors.Common.BadRequest
	}

	task, err := h.service.GetTaskByID(ctx, taskID, performerID)
	if err != nil {
		return nil, err
	}

	if task.Status == status {
		ctx.Logger().Text("task status does not change, respond")
		return &dto.ChangeTaskStatusResponse{}, nil
	}

	goal, err := h.service.GetGoalByID(ctx, task.GoalID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("change task completed status")
	task.SetStatus(status)

	ctx.Logger().Text("update task in db")
	if err = h.taskRepository.Update(ctx, *task); err != nil {
		ctx.Logger().Error("failed to update task in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("adjust goal stats")
	adjustValue := 1
	if !task.IsDone() {
		adjustValue = -1
	}
	goal.AdjustTotalDoneTask(adjustValue)

	ctx.Logger().Text("update goal in db")
	if err = h.goalRepository.Update(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to update goal in db", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done change task completed status request")
	return &dto.ChangeTaskStatusResponse{}, nil
}
