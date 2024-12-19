package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type ChangeHabitStatusHandler struct {
	habitRepository           domain.HabitRepository
	habitDailyStatsRepository domain.HabitDailyStatsRepository
	service                   domain.Service
}

func NewChangeHabitStatusHandler(
	habitRepository domain.HabitRepository,
	habitDailyStatsRepository domain.HabitDailyStatsRepository,
	service domain.Service,
) ChangeHabitStatusHandler {
	return ChangeHabitStatusHandler{
		habitRepository:           habitRepository,
		habitDailyStatsRepository: habitDailyStatsRepository,
		service:                   service,
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

	if habit.IsActive() {
		ctx.Logger().Text("update today habit stats")
		dailyStats, dErr := h.habitDailyStatsRepository.FindByDate(ctx, performerID, manipulation.NowUTC())
		if dErr != nil {
			ctx.Logger().Error("failed to update today habit stats", dErr, appcontext.Fields{})
		}
		if dailyStats != nil {
			dailyStats.AddNewHabit(habit.ID)

			ctx.Logger().Text("update today habit stats in db")
			if err = h.habitDailyStatsRepository.Update(ctx, *dailyStats); err != nil {
				ctx.Logger().Error("failed to update today habit stats in db", err, appcontext.Fields{})
				return nil, err
			}
		}
	}

	ctx.Logger().Text("done change habit status request")
	return &dto.ChangeHabitStatusResponse{}, nil
}
