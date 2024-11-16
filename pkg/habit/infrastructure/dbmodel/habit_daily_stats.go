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
	UserID         primitive.ObjectID `bson:"userId"`
	Date           time.Time          `bson:"date"`
	ScheduledCount int                `bson:"scheduledCount"`
	CompletedCount int                `bson:"completedCount"`
	CompletedIDs   []string           `bson:"completedIds"`
}

func (s HabitDailyStats) ToDomain() domain.HabitDailyStats {
	return domain.HabitDailyStats{
		ID:             s.ID.Hex(),
		UserID:         s.UserID.Hex(),
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

	uid, err := database.ObjectIDFromString(stats.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitDailyStats{
		ID:             id,
		UserID:         uid,
		Date:           stats.Date,
		ScheduledCount: stats.ScheduledCount,
		CompletedCount: stats.CompletedCount,
		CompletedIDs:   stats.CompletedIDs,
	}, nil
}
