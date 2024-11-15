package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type UpdateHabitHandler struct {
	habitRepository domain.HabitRepository
	service         domain.Service
}

func NewUpdateHabitHandler(habitRepository domain.HabitRepository, service domain.Service) UpdateHabitHandler {
	return UpdateHabitHandler{
		habitRepository: habitRepository,
		service:         service,
	}
}

func (h UpdateHabitHandler) UpdateHabit(ctx *appcontext.AppContext, performerID, habitID string, req dto.UpdateHabitRequest) (*dto.UpdateHabitResponse, error) {
	ctx.Logger().Info("new update habit request", appcontext.Fields{
		"performerID": performerID, "name": req.Name, "goal": req.Goal,
		"daysOfWeek": req.DayOfWeeks, "icon": req.Icon, "sortOrder": req.SortOrder,
	})

	habit, err := h.service.GetHabitByID(ctx, habitID, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("set habit data")
	if err = habit.SetName(req.Name); err != nil {
		ctx.Logger().Error("failed to set habit name", err, appcontext.Fields{})
		return nil, err
	}
	if err = habit.SetGoal(req.Goal); err != nil {
		ctx.Logger().Error("failed to set habit goal", err, appcontext.Fields{})
		return nil, err
	}
	if err = habit.SetDaysOfWeek(req.DayOfWeeks); err != nil {
		ctx.Logger().Error("failed to set habit days of week", err, appcontext.Fields{})
		return nil, err
	}
	if err = habit.SetIcon(req.Icon); err != nil {
		ctx.Logger().Error("failed to set habit icon", err, appcontext.Fields{})
		return nil, err
	}
	habit.SetSortOrder(req.SortOrder)

	ctx.Logger().Text("update habit in db")
	if err = h.habitRepository.Update(ctx, *habit); err != nil {
		ctx.Logger().Error("failed to update habit in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update habit request")
	return &dto.UpdateHabitResponse{
		ID: habit.ID,
	}, nil
}
