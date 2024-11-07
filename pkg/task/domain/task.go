package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type TaskRepository interface {
	Create(ctx *appcontext.AppContext, task Task) error
	Update(ctx *appcontext.AppContext, task Task) error
	FindByFilter(ctx *appcontext.AppContext, filter TaskFilter) ([]Task, error)
}

type Task struct {
	ID           string
	UserID       string
	GoalID       string
	Name         string
	Description  string
	SearchString string
	DueDate      *time.Time
	IsCompleted  bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
}
