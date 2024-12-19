package domain

import (
	"slices"
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
	FindByDate(ctx *appcontext.AppContext, userID string, date time.Time) (*HabitDailyStats, error)
	FindByFilter(ctx *appcontext.AppContext, filter HabitDailyStatsFilter) ([]HabitDailyStats, error)
}

const (
	StatsDefaultPreviousDays = 6
)

type HabitDailyStats struct {
	ID           string
	UserID       string
	Date         time.Time
	IsCompleted  bool
	ScheduledIDs []string
	CompletedIDs []string
}

func NewHabitDailyStats(userID string, scheduledIDs []string, date time.Time) (*HabitDailyStats, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitDailyStats{
		ID:           database.NewStringID(),
		UserID:       userID,
		Date:         date,
		IsCompleted:  false,
		ScheduledIDs: scheduledIDs,
		CompletedIDs: make([]string, 0),
	}, nil
}

func (s *HabitDailyStats) HabitCompleted(habitID string) {
	s.CompletedIDs = append(s.CompletedIDs, habitID)

	if len(s.CompletedIDs) == len(s.ScheduledIDs) {
		s.IsCompleted = true
	}
}

func (s *HabitDailyStats) AddNewHabit(habitID string) {
	if slices.Contains(s.ScheduledIDs, habitID) {
		s.ScheduledIDs = append(s.ScheduledIDs, habitID)
	}
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
		FromDate: fromDate,
	}, nil
}
