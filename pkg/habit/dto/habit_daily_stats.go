package dto

import (
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

type HabitDailyStats struct {
	ID             string                    `json:"id"`
	Date           *httprespond.TimeResponse `json:"date"`
	ScheduledCount int                       `json:"scheduledCount"`
	CompletedCount int                       `json:"completedCount"`
	CompletedIDs   []string                  `json:"completedIDs"`
}

func (HabitDailyStats) FromDomain(stats domain.HabitDailyStats) HabitDailyStats {
	return HabitDailyStats{
		ID:             stats.ID,
		Date:           httprespond.NewTimeResponse(stats.Date),
		ScheduledCount: stats.ScheduledCount,
		CompletedCount: stats.CompletedCount,
		CompletedIDs:   stats.CompletedIDs,
	}
}
