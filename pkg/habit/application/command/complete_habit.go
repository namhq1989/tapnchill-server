package command

import (
	"slices"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type CompleteHabitHandler struct {
	habitRepository           domain.HabitRepository
	habitCompletionRepository domain.HabitCompletionRepository
	habitDailyStatsRepository domain.HabitDailyStatsRepository
	service                   domain.Service
}

func NewCompleteHabitHandler(
	habitRepository domain.HabitRepository,
	habitCompletionRepository domain.HabitCompletionRepository,
	habitDailyStatsRepository domain.HabitDailyStatsRepository,
	service domain.Service,
) CompleteHabitHandler {
	return CompleteHabitHandler{
		habitRepository:           habitRepository,
		habitCompletionRepository: habitCompletionRepository,
		habitDailyStatsRepository: habitDailyStatsRepository,
		service:                   service,
	}
}

func (h CompleteHabitHandler) CompleteHabit(ctx *appcontext.AppContext, performerID, habitID string, req dto.CompleteHabitRequest) (*dto.CompleteHabitResponse, error) {
	ctx.Logger().Info("new complete habit request", appcontext.Fields{"performerID": performerID, "date": req.Date})

	habit, err := h.service.GetHabitByID(ctx, habitID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("get start of day with client date")
	startOfDay, err := manipulation.GetStartOfDayWithClientDate(req.Date)
	if err != nil {
		ctx.Logger().Error("failed to get start of day with client date", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find daily stats in db")
	stats, err := h.habitDailyStatsRepository.FindByDate(ctx, performerID, *startOfDay)
	if err != nil {
		ctx.Logger().Error("failed to find daily stats in db", err, appcontext.Fields{})
		return nil, err
	}

	if stats == nil {
		ctx.Logger().Text("not found, find user habits to create a new daily stats")
		habits, hErr := h.service.GetUserHabits(ctx, performerID)
		if hErr != nil {
			ctx.Logger().Error("failed to find user habits to create a new daily stats", hErr, appcontext.Fields{})
			return nil, hErr
		}

		ctx.Logger().Text("find scheduled habits of today")
		scheduledIDs := make([]string, 0)
		for _, hb := range habits {
			if hb.IsActive() && hb.CreatedAt.Before(*startOfDay) && slices.Contains(hb.DaysOfWeek, int(startOfDay.Weekday())) {
				scheduledIDs = append(scheduledIDs, hb.ID)
			}
		}

		ctx.Logger().Text("create new daily stats model")
		stats, err = domain.NewHabitDailyStats(performerID, scheduledIDs, *startOfDay)
		if err != nil {
			ctx.Logger().Error("failed to create new daily stats model", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("persist daily stats in db")
		if err = h.habitDailyStatsRepository.Create(ctx, *stats); err != nil {
			ctx.Logger().Error("failed to persist daily stats in db", err, appcontext.Fields{})
			return nil, err
		}
	} else if slices.Contains(stats.CompletedIDs, habitID) {
		ctx.Logger().Text("already completed, respond")
		return &dto.CompleteHabitResponse{}, nil
	}

	ctx.Logger().Text("create new habit completion model")
	completion, err := domain.NewHabitCompletion(habitID, *startOfDay)
	if err != nil {
		ctx.Logger().Error("failed to create new habit completion model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist habit completion in db")
	if err = h.habitCompletionRepository.Create(ctx, *completion); err != nil {
		ctx.Logger().Error("failed to persist habit completion in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update daily stats")
	stats.HabitCompleted(habitID)

	ctx.Logger().Text("update daily stats in db")
	if err = h.habitDailyStatsRepository.Update(ctx, *stats); err != nil {
		ctx.Logger().Error("failed to update daily stats in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update habit stats")
	habit.OnCompleted(*startOfDay)

	ctx.Logger().Text("update habit in db")
	if err = h.habitRepository.Update(ctx, *habit); err != nil {
		ctx.Logger().Error("failed to update habit in db", err, appcontext.Fields{})
		return nil, err
	}

	_ = h.service.DeleteUserCaching(ctx, performerID, manipulation.NowUTC())

	ctx.Logger().Text("done complete habit request")
	return &dto.CompleteHabitResponse{}, nil
}
