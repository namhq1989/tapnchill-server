package shared

import (
	"fmt"
	"sort"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
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
	fromDate := startOfDay.AddDate(0, 0, -days).UTC()

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

	ctx.Logger().Info("find in db", appcontext.Fields{"filter": filter})
	stats, err = s.habitDailyStatsRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	if len(stats) < days+1 { // +1 because we need to count today's stats in
		ctx.Logger().Text("fetching user habits")
		habits, hErr := s.GetUserHabits(ctx, userID)
		if hErr != nil {
			ctx.Logger().Error("failed to fetch user habits", hErr, appcontext.Fields{})
			return nil, hErr
		}

		ctx.Logger().Text("generating default stats for missing dates")
		stats = s.generateDefaultStatsIfMissing(ctx, stats, fromDate, startOfDay.UTC(), habits)
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Date.After(stats[j].Date)
	})

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetUserStats(ctx, userID, dateDDMM, stats); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
	}

	return stats, nil
}

func (s Service) generateDefaultStatsIfMissing(ctx *appcontext.AppContext, stats []domain.HabitDailyStats, fromDate time.Time, startOfDay time.Time, habits []domain.Habit) []domain.HabitDailyStats {
	for d := fromDate; !d.After(startOfDay); d = d.AddDate(0, 0, 1) {
		isExisted := false
		for _, stat := range stats {
			if stat.Date.YearDay() == d.YearDay() {
				isExisted = true
				break
			}
		}

		if isExisted {
			continue
		}

		stats = append(stats, domain.HabitDailyStats{
			ID:           database.NewStringID(),
			Date:         d,
			IsCompleted:  false,
			ScheduledIDs: s.getScheduledIDsForDay(ctx, d, habits),
			CompletedIDs: []string{},
		})
	}

	return stats
}

func (Service) getScheduledIDsForDay(ctx *appcontext.AppContext, date time.Time, habits []domain.Habit) []string {
	ctx.Logger().Text("calculating scheduled count")

	dayOfWeek := int(date.Weekday())

	ids := make([]string, 0)
	for _, habit := range habits {
		for _, scheduledDay := range habit.DaysOfWeek {
			if scheduledDay == dayOfWeek && habit.IsActive() {
				ids = append(ids, habit.ID)
				break
			}
		}
	}

	ctx.Logger().Info("scheduled ids calculated", appcontext.Fields{"ids": ids})
	return ids
}
