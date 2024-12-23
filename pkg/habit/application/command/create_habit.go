package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type CreateHabitHandler struct {
	habitRepository           domain.HabitRepository
	habitDailyStatsRepository domain.HabitDailyStatsRepository
	service                   domain.Service
	userHub                   domain.UserHub
}

func NewCreateHabitHandler(
	habitRepository domain.HabitRepository,
	habitDailyStatsRepository domain.HabitDailyStatsRepository,
	service domain.Service,
	userHub domain.UserHub,
) CreateHabitHandler {
	return CreateHabitHandler{
		habitRepository:           habitRepository,
		habitDailyStatsRepository: habitDailyStatsRepository,
		service:                   service,
		userHub:                   userHub,
	}
}

func (h CreateHabitHandler) CreateHabit(ctx *appcontext.AppContext, performerID string, req dto.CreateHabitRequest) (*dto.CreateHabitResponse, error) {
	ctx.Logger().Info("new create habit request", appcontext.Fields{
		"performerID": performerID, "date": req.Date, "name": req.Name, "goal": req.Goal,
		"daysOfWeek": req.DaysOfWeek, "icon": req.Icon, "sortOrder": req.SortOrder,
	})

	ctx.Logger().Text("get user habit quota")
	quota, isFreePlan, err := h.userHub.GetHabitQuota(ctx, performerID)
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
		err = apperrors.User.FreePlanLimitReached
		if !isFreePlan {
			err = apperrors.User.ProPlanLimitReached
		}

		return nil, err
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

	ctx.Logger().Text("update today habit stats")
	dailyStats, err := h.habitDailyStatsRepository.FindByDate(ctx, performerID, manipulation.NowUTC())
	if err != nil {
		ctx.Logger().Error("failed to update today habit stats", err, appcontext.Fields{})
	}
	if dailyStats != nil {
		dailyStats.AddNewHabit(habit.ID)

		ctx.Logger().Text("update today habit stats in db")
		if err = h.habitDailyStatsRepository.Update(ctx, *dailyStats); err != nil {
			ctx.Logger().Error("failed to update today habit stats in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("done create habit request")
	return &dto.CreateHabitResponse{
		ID: habit.ID,
	}, nil
}
