package dto

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
)

type Task struct {
	ID          string                    `json:"id"`
	GoalID      string                    `json:"goalId"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	DueDate     *httprespond.TimeResponse `json:"dueDate"`
	Status      string                    `json:"status"`
	CreatedAt   *httprespond.TimeResponse `json:"createdAt"`
	CompletedAt *httprespond.TimeResponse `json:"completedAt"`
}

func (Task) FromDomain(task domain.Task) Task {
	t := Task{
		ID:          task.ID,
		GoalID:      task.GoalID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status.String(),
		CreatedAt:   httprespond.NewTimeResponse(task.CreatedAt),
	}

	if task.DueDate != nil {
		t.DueDate = httprespond.NewTimeResponse(*task.DueDate)
	} else {
		t.DueDate = httprespond.NewTimeResponse(time.Time{})
	}

	if task.CompletedAt != nil {
		t.CompletedAt = httprespond.NewTimeResponse(*task.CompletedAt)
	} else {
		t.CompletedAt = httprespond.NewTimeResponse(time.Time{})
	}

	return t
}
