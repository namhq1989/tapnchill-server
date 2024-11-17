package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type CreateHabitHandler struct {
	habitRepository domain.HabitRepository
	service         domain.Service
}

func NewCreateHabitHandler(habitRepository domain.HabitRepository, service domain.Service) CreateHabitHandler {
	return CreateHabitHandler{
		habitRepository: habitRepository,
		service:         service,
	}
}

func (h CreateHabitHandler) CreateHabit(ctx *appcontext.AppContext, performerID string, req dto.CreateHabitRequest) (*dto.CreateHabitResponse, error) {
	ctx.Logger().Info("new create habit request", appcontext.Fields{
		"performerID": performerID, "date": req.Date, "name": req.Name, "goal": req.Goal,
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

	_ = h.service.DeleteUserCaching(ctx, performerID, req.Date)

	ctx.Logger().Text("done create habit request")
	return &dto.CreateHabitResponse{
		ID: habit.ID,
	}, nil
}
