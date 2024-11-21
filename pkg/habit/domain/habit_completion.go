package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
)

type HabitCompletionRepository interface {
	Create(ctx *appcontext.AppContext, completion HabitCompletion) error
}

type HabitCompletion struct {
	ID          string
	HabitID     string
	CompletedAt time.Time
}

func NewHabitCompletion(habitID string, date time.Time) (*HabitCompletion, error) {
	return &HabitCompletion{
		ID:          database.NewStringID(),
		HabitID:     habitID,
		CompletedAt: date,
	}, nil
}
