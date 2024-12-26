package domain

import (
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type TaskRepository interface {
	Create(ctx *appcontext.AppContext, task Task) error
	Update(ctx *appcontext.AppContext, task Task) error
	Delete(ctx *appcontext.AppContext, taskID string) error
	CountByGoalID(ctx *appcontext.AppContext, goalID string) (int64, error)
	FindByFilter(ctx *appcontext.AppContext, filter TaskFilter) ([]Task, error)
	FindByID(ctx *appcontext.AppContext, taskID string) (*Task, error)
}

type Task struct {
	ID           string
	UserID       string
	GoalID       string
	Name         string
	Description  string
	SearchString string
	DueDate      *time.Time
	Status       TaskStatus
	CreatedAt    time.Time
	CompletedAt  *time.Time
}

func NewTask(userID, goalID, name, description string, dueDate *time.Time) (*Task, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if !database.IsValidObjectID(goalID) {
		return nil, apperrors.Task.InvalidGoalID
	}

	t := &Task{
		ID:           database.NewStringID(),
		UserID:       userID,
		GoalID:       goalID,
		Name:         name,
		Description:  description,
		SearchString: manipulation.NormalizeText(fmt.Sprintf("%s %s", name, description)),
		DueDate:      dueDate,
		Status:       TaskStatusTodo,
		CreatedAt:    manipulation.NowUTC(),
	}

	if err := t.SetName(name); err != nil {
		return nil, err
	}

	if err := t.SetDescription(description); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Task) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	t.Name = name
	t.SearchString = manipulation.NormalizeText(fmt.Sprintf("%s", name))
	return nil
}

func (t *Task) SetDescription(description string) error {
	if len(description) > 1000 {
		return apperrors.Common.InvalidDescription
	}

	t.Description = description
	return nil
}

func (t *Task) SetDueDate(dueDate *time.Time) {
	t.DueDate = dueDate
}

func (t *Task) SetStatus(status TaskStatus) {
	t.Status = status

	if status == TaskStatusTodo {
		t.CompletedAt = nil
	} else {
		now := manipulation.NowUTC()
		t.CompletedAt = &now
	}
}

func (t *Task) IsDone() bool {
	return t.Status == TaskStatusDone
}
