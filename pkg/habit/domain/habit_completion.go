package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type HabitCompletion struct {
	ID          string
	HabitID     string
	CompletedAt time.Time
}

func NewHabitCompletion(habitID string) (*HabitCompletion, error) {
	return &HabitCompletion{
		ID:          database.NewStringID(),
		HabitID:     habitID,
		CompletedAt: manipulation.NowUTC(),
	}, nil
}
