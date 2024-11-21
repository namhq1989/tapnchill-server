package dto

import (
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

type HabitDailyStats struct {
	ID           string                    `json:"id"`
	Date         *httprespond.TimeResponse `json:"date"`
	IsCompleted  bool                      `json:"isCompleted"`
	ScheduledIDs []string                  `json:"scheduledIds"`
	CompletedIDs []string                  `json:"completedIds"`
}

func (HabitDailyStats) FromDomain(stats domain.HabitDailyStats) HabitDailyStats {
	return HabitDailyStats{
		ID:           stats.ID,
		Date:         httprespond.NewTimeResponse(stats.Date),
		IsCompleted:  stats.IsCompleted,
		ScheduledIDs: stats.ScheduledIDs,
		CompletedIDs: stats.CompletedIDs,
	}
}
