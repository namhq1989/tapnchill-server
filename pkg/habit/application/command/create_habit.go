package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type CreateHabitHandler struct {
	habitRepository   domain.HabitRepository
	cachingRepository domain.CachingRepository
}

func NewCreateHabitHandler(habitRepository domain.HabitRepository, cachingRepository domain.CachingRepository) CreateHabitHandler {
	return CreateHabitHandler{
		habitRepository:   habitRepository,
		cachingRepository: cachingRepository,
	}
}

func (h CreateHabitHandler) CreateHabit(ctx *appcontext.AppContext, performerID string, req dto.CreateHabitRequest) (*dto.CreateHabitResponse, error) {
	ctx.Logger().Info("new create habit request", appcontext.Fields{
		"performerID": performerID, "name": req.Name, "goal": req.Goal,
		"daysOfWeek": req.DaysOfWeek, "icon": req.Icon, "sortOrder": req.SortOrder,
	})

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

	ctx.Logger().Text("delete in caching")
	if err = h.cachingRepository.DeleteUserHabits(ctx, performerID); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done create habit request")
	return &dto.CreateHabitResponse{
		ID: habit.ID,
	}, nil
}
