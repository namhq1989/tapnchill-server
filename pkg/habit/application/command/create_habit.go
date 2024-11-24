package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type CreateHabitHandler struct {
	habitRepository domain.HabitRepository
	service         domain.Service
	userHub         domain.UserHub
}

func NewCreateHabitHandler(habitRepository domain.HabitRepository, service domain.Service, userHub domain.UserHub) CreateHabitHandler {
	return CreateHabitHandler{
		habitRepository: habitRepository,
		service:         service,
		userHub:         userHub,
	}
}

func (h CreateHabitHandler) CreateHabit(ctx *appcontext.AppContext, performerID string, req dto.CreateHabitRequest) (*dto.CreateHabitResponse, error) {
	ctx.Logger().Info("new create habit request", appcontext.Fields{
		"performerID": performerID, "date": req.Date, "name": req.Name, "goal": req.Goal,
		"daysOfWeek": req.DaysOfWeek, "icon": req.Icon, "sortOrder": req.SortOrder,
	})

	ctx.Logger().Text("get user habit quota")
	quota, err := h.userHub.GetHabitQuota(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get user habit quota", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count user total habits")
	totalHabits, err := h.habitRepository.CountByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to count user total habits", err, appcontext.Fields{})
		return nil, err
	}

	if totalHabits >= quota {
		ctx.Logger().Error("user habit quota exceeded", err, appcontext.Fields{"quota": quota, "total": totalHabits})
		return nil, apperrors.User.ResourceLimitReached
	}

	ctx.Logger().Text("create new habit model")
	habit, err := domain.NewHabit(performerID, req.Name, req.Goal, req.DaysOfWeek, req.Icon, req.SortOrder)
	if err != nil {
		ctx.Logger().Error("failed to create new habit model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist habit in db")
	if err = h.habitRepository.Create(ctx, *habit); err != nil {
		ctx.Logger().Error("failed to persist habit in db", err, appcontext.Fields{})
		return nil, err
	}

	_ = h.service.DeleteUserCaching(ctx, performerID, manipulation.NowUTC())

	ctx.Logger().Text("done create habit request")
	return &dto.CreateHabitResponse{
		ID: habit.ID,
	}, nil
}
