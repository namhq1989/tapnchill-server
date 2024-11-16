package shared

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

func (s Service) GetUserStats(ctx *appcontext.AppContext, userID, date string, days int) ([]domain.HabitDailyStats, error) {
	ctx.Logger().Info("get user stats", appcontext.Fields{"userID": userID, "days": days})

	ctx.Logger().Text("calculate the from date")
	startOfDay, err := manipulation.GetStartOfDayWithClientDate(date)
	if err != nil {
		ctx.Logger().Error("failed to calculate the from date", err, appcontext.Fields{})
		return nil, err
	}
	dateDDMM := fmt.Sprintf("%02d%02d", startOfDay.Day(), int(startOfDay.Month()))
	fromDate := startOfDay.AddDate(0, 0, -days)

	ctx.Logger().Text("find in caching")
	stats, err := s.cachingRepository.GetUserStats(ctx, userID, dateDDMM)
	if stats != nil {
		ctx.Logger().Text("found in caching, respond")
		return stats, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("create new filter")
	filter, err := domain.NewHabitDailyStatsFilter(userID, fromDate)
	if err != nil {
		ctx.Logger().Error("failed to create new filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find in db")
	stats, err = s.habitDailyStatsRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetUserStats(ctx, userID, dateDDMM, stats); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
	}

	return stats, nil
}
