package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HabitDailyStats struct {
	ID             primitive.ObjectID `bson:"_id"`
	HabitID        primitive.ObjectID `bson:"habitId"`
	Date           time.Time          `bson:"date"`
	ScheduledCount int                `bson:"scheduledCount"`
	CompletedCount int                `bson:"completedCount"`
	CompletedIDs   []string           `bson:"completedIds"`
}

func (s HabitDailyStats) ToDomain() domain.HabitDailyStats {
	return domain.HabitDailyStats{
		ID:             s.ID.Hex(),
		HabitID:        s.HabitID.Hex(),
		Date:           s.Date,
		ScheduledCount: s.ScheduledCount,
		CompletedCount: s.CompletedCount,
		CompletedIDs:   s.CompletedIDs,
	}
}

func (HabitDailyStats) FromDomain(stats domain.HabitDailyStats) (*HabitDailyStats, error) {
	id, err := database.ObjectIDFromString(stats.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	hid, err := database.ObjectIDFromString(stats.HabitID)
	if err != nil {
		return nil, apperrors.Habit.InvalidID
	}

	return &HabitDailyStats{
		ID:             id,
		HabitID:        hid,
		Date:           stats.Date,
		ScheduledCount: stats.ScheduledCount,
		CompletedCount: stats.CompletedCount,
		CompletedIDs:   stats.CompletedIDs,
	}, nil
}
