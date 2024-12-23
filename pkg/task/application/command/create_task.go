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
	service        domain.Service
	userHub        domain.UserHub
}

func NewCreateTaskHandler(
	taskRepository domain.TaskRepository,
	goalRepository domain.GoalRepository,
	service domain.Service,
	userHub domain.UserHub,
) CreateTaskHandler {
	return CreateTaskHandler{
		taskRepository: taskRepository,
		goalRepository: goalRepository,
		service:        service,
		userHub:        userHub,
	}
}

func (h CreateTaskHandler) CreateTask(ctx *appcontext.AppContext, performerID string, req dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	ctx.Logger().Info("new create task request", appcontext.Fields{
		"performerID": performerID, "goalID": req.GoalID,
		"name": req.Name, "description": req.Description, "dueDate": req.DueDate,
	})

	goal, err := h.service.GetGoalByID(ctx, req.GoalID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("get user task quota")
	quota, isFreePlan, err := h.userHub.GetTaskQuota(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get user task quota", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count goal total tasks")
	totalTasks, err := h.taskRepository.CountByGoalID(ctx, req.GoalID)
	if err != nil {
		ctx.Logger().Error("failed to count user total tasks", err, appcontext.Fields{})
		return nil, err
	}

	if totalTasks >= quota {
		ctx.Logger().Error("user task quota exceeded", err, appcontext.Fields{"quota": quota, "total": totalTasks})
		err = apperrors.User.FreePlanLimitReached
		if !isFreePlan {
			err = apperrors.User.ProPlanLimitReached
		}

		return nil, err
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

	ctx.Logger().Text("adjust goal stats")
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
