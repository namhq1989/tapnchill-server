package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

type HabitDailyStatsRepository interface {
	Create(ctx *appcontext.AppContext, stats HabitDailyStats) error
	Update(ctx *appcontext.AppContext, stats HabitDailyStats) error
	FindByID(ctx *appcontext.AppContext, statsID string) (*HabitDailyStats, error)
	FindByFilter(ctx *appcontext.AppContext, filter HabitDailyStatsFilter) ([]HabitDailyStats, error)
}

type HabitDailyStats struct {
	ID             string
	HabitID        string
	Date           time.Time
	ScheduledCount int
	CompletedCount int
	CompletedIDs   []string
}

func NewHabitDailyStats(habitID string, date time.Time) (*HabitDailyStats, error) {
	if !database.IsValidObjectID(habitID) {
		return nil, apperrors.Habit.InvalidID
	}

	return &HabitDailyStats{
		ID:             database.NewStringID(),
		HabitID:        habitID,
		Date:           date,
		ScheduledCount: 0,
		CompletedCount: 0,
		CompletedIDs:   make([]string, 0),
	}, nil
}

func (s *HabitDailyStats) SetScheduledCount(count int) error {
	s.ScheduledCount = count
	return nil
}

func (s *HabitDailyStats) HabitCompleted(habitID string) error {
	s.CompletedCount += 1
	s.CompletedIDs = append(s.CompletedIDs, habitID)
	return nil
}

type HabitDailyStatsFilter struct {
	HabitID  string
	FromDate time.Time
}

func NewHabitDailyStatsFilter(habitID string, fromDate time.Time) (*HabitDailyStatsFilter, error) {
	if !database.IsValidObjectID(habitID) {
		return nil, apperrors.Habit.InvalidID
	}

	return &HabitDailyStatsFilter{
		HabitID:  habitID,
		FromDate: fromDate,
	}, nil
}
