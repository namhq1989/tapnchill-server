package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HabitCompletion struct {
	ID          primitive.ObjectID `bson:"_id"`
	HabitID     primitive.ObjectID `bson:"habitId"`
	CompletedAt time.Time          `bson:"completedAt"`
}

func (c HabitCompletion) ToDomain() domain.HabitCompletion {
	return domain.HabitCompletion{
		ID:          c.ID.Hex(),
		HabitID:     c.HabitID.Hex(),
		CompletedAt: c.CompletedAt,
	}
}

func (HabitCompletion) FromDomain(completion domain.HabitCompletion) (*HabitCompletion, error) {
	id, err := database.ObjectIDFromString(completion.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	hid, err := database.ObjectIDFromString(completion.HabitID)
	if err != nil {
		return nil, apperrors.Habit.InvalidID
	}

	return &HabitCompletion{
		ID:          id,
		HabitID:     hid,
		CompletedAt: completion.CompletedAt,
	}, nil
}
