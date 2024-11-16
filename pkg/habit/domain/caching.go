package domain

import "github.com/namhq1989/go-utilities/appcontext"

type CachingRepository interface {
	GetUserHabits(ctx *appcontext.AppContext, userID string) ([]Habit, error)
	SetUserHabits(ctx *appcontext.AppContext, userID string, habits []Habit) error
	DeleteUserHabits(ctx *appcontext.AppContext, userID string) error

	GetUserStats(ctx *appcontext.AppContext, userID, date string) ([]HabitDailyStats, error)
	SetUserStats(ctx *appcontext.AppContext, userID, date string, stats []HabitDailyStats) error
	DeleteUserStats(ctx *appcontext.AppContext, userID, date string) error
}
