package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

type HabitDailyStats struct {
	ID             string
	UserID         string
	Date           time.Time
	ScheduledCount int
	CompletedCount int
}

func NewHabitDailyStats(userID string, date time.Time) (*HabitDailyStats, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitDailyStats{
		ID:             database.NewStringID(),
		UserID:         userID,
		Date:           date,
		ScheduledCount: 0,
		CompletedCount: 0,
	}, nil
}

func (s *HabitDailyStats) SetScheduledCount(count int) error {
	s.ScheduledCount = count
	return nil
}

func (s *HabitDailyStats) IncreaseCompletedCount(value int) error {
	s.CompletedCount += value
	return nil
}
