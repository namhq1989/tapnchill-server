package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetHabitByID(ctx *appcontext.AppContext, habitID, userID string) (*Habit, error)
}
