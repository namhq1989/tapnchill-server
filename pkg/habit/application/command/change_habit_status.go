package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type ChangeHabitStatusHandler struct {
	habitRepository   domain.HabitRepository
	cachingRepository domain.CachingRepository
	service           domain.Service
}

func NewChangeHabitStatusHandler(habitRepository domain.HabitRepository, cachingRepository domain.CachingRepository, service domain.Service) ChangeHabitStatusHandler {
	return ChangeHabitStatusHandler{
		habitRepository:   habitRepository,
		cachingRepository: cachingRepository,
		service:           service,
	}
}

func (h ChangeHabitStatusHandler) ChangeHabitStatus(ctx *appcontext.AppContext, performerID, habitID string, req dto.ChangeHabitStatusRequest) (*dto.ChangeHabitStatusResponse, error) {
	ctx.Logger().Info("new change habit status request", appcontext.Fields{
		"performerID": performerID, "status": req.Status,
	})

	ctx.Logger().Text("check status")
	status := domain.ToHabitStatus(req.Status)
	if !status.IsValid() {
		ctx.Logger().ErrorText("invalid status, respond")
		return nil, apperrors.Common.BadRequest
	}

	habit, err := h.service.GetHabitByID(ctx, habitID, performerID)
	if err != nil {
		return nil, err
	}

	if habit.Status == status {
		ctx.Logger().Text("habit status does not change, respond")
		return &dto.ChangeHabitStatusResponse{}, nil
	}

	ctx.Logger().Text("change habit status")
	habit.SetStatus(status)

	ctx.Logger().Text("update habit in db")
	if err = h.habitRepository.Update(ctx, *habit); err != nil {
		ctx.Logger().Error("failed to update habit in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("delete in caching")
	if err = h.cachingRepository.DeleteUserHabits(ctx, performerID); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done change habit status request")
	return &dto.ChangeHabitStatusResponse{}, nil
}
