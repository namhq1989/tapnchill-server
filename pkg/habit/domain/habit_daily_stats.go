package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HabitDailyStatsRepository interface {
	Create(ctx *appcontext.AppContext, stats HabitDailyStats) error
	Update(ctx *appcontext.AppContext, stats HabitDailyStats) error
	FindByID(ctx *appcontext.AppContext, statsID string) (*HabitDailyStats, error)
	FindByDate(ctx *appcontext.AppContext, habitID string, date time.Time) (*HabitDailyStats, error)
	FindByFilter(ctx *appcontext.AppContext, filter HabitDailyStatsFilter) ([]HabitDailyStats, error)
}

const (
	StatsDefaultPreviousDays = 5
)

type HabitDailyStats struct {
	ID             string
	UserID         string
	Date           time.Time
	ScheduledCount int
	CompletedCount int
	CompletedIDs   []string
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
		CompletedIDs:   make([]string, 0),
	}, nil
}

func (s *HabitDailyStats) SetScheduledCount(count int) {
	s.ScheduledCount = count
}

func (s *HabitDailyStats) HabitCompleted(habitID string) {
	s.CompletedCount += 1
	s.CompletedIDs = append(s.CompletedIDs, habitID)
}

type HabitDailyStatsFilter struct {
	UserID   primitive.ObjectID
	FromDate time.Time
}

func NewHabitDailyStatsFilter(userID string, fromDate time.Time) (*HabitDailyStatsFilter, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitDailyStatsFilter{
		UserID:   uid,
		FromDate: fromDate.UTC(),
	}, nil
}
