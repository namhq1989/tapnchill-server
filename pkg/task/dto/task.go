package dto

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
)

type Task struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	DueDate     *httprespond.TimeResponse `json:"dueDate"`
	IsCompleted bool                      `json:"isCompleted"`
	CreatedAt   time.Time                 `json:"createdAt"`
	CompletedAt *httprespond.TimeResponse `json:"completedAt"`
}
