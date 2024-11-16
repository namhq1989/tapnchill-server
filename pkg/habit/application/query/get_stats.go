package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type GetStatsHandler struct {
	service domain.Service
}

func NewGetStatsHandler(service domain.Service) GetStatsHandler {
	return GetStatsHandler{
		service: service,
	}
}

func (h GetStatsHandler) GetStats(ctx *appcontext.AppContext, performerID string, req dto.GetStatsRequest) (*dto.GetStatsResponse, error) {
	ctx.Logger().Info("new get stats request", appcontext.Fields{"performerID": performerID, "date": req.Date})

	stats, err := h.service.GetUserStats(ctx, performerID, req.Date, domain.StatsDefaultPreviousDays)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.HabitDailyStats, 0)
	for _, stat := range stats {
		result = append(result, dto.HabitDailyStats{}.FromDomain(stat))
	}

	ctx.Logger().Text("done get stats request")
	return &dto.GetStatsResponse{
		Stats: result,
	}, nil
}
