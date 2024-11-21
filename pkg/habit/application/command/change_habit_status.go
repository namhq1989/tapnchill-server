package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type ChangeHabitStatusHandler struct {
	habitRepository domain.HabitRepository
	service         domain.Service
}

func NewChangeHabitStatusHandler(habitRepository domain.HabitRepository, service domain.Service) ChangeHabitStatusHandler {
	return ChangeHabitStatusHandler{
		habitRepository: habitRepository,
		service:         service,
	}
}

func (h ChangeHabitStatusHandler) ChangeHabitStatus(ctx *appcontext.AppContext, performerID, habitID string, req dto.ChangeHabitStatusRequest) (*dto.ChangeHabitStatusResponse, error) {
	ctx.Logger().Info("new change habit status request", appcontext.Fields{
		"performerID": performerID, "date": req.Date, "status": req.Status,
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

	_ = h.service.DeleteUserCaching(ctx, performerID, manipulation.NowUTC())

	ctx.Logger().Text("done change habit status request")
	return &dto.ChangeHabitStatusResponse{}, nil
}
