package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type Service interface {
	GetHabitByID(ctx *appcontext.AppContext, habitID, userID string) (*Habit, error)
	GetUserHabits(ctx *appcontext.AppContext, userID string) ([]Habit, error)
	GetUserStats(ctx *appcontext.AppContext, userID, date string, days int) ([]HabitDailyStats, error)
	DeleteUserCaching(ctx *appcontext.AppContext, userID string, date time.Time) error
}
