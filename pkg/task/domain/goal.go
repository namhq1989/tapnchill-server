package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type GoalRepository interface {
	Create(ctx *appcontext.AppContext, goal Goal) error
	Update(ctx *appcontext.AppContext, goal Goal) error
	FindByFilter(ctx *appcontext.AppContext, filter GoalFilter) ([]Goal, error)
}

type Goal struct {
	ID           string
	UserID       string
	Name         string
	Description  string
	SearchString string
	Stats        GoalStats
	IsCompleted  bool
	CreatedAt    time.Time
}

type GoalStats struct {
	TotalTask          int
	TotalCompletedTask int
}
