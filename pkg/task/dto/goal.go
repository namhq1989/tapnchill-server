package dto

import (
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
)

type Goal struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Stats       GoalStats                 `json:"stats"`
	IsCompleted bool                      `json:"isCompleted"`
	CreatedAt   *httprespond.TimeResponse `json:"createdAt"`
}

type GoalStats struct {
	TotalTask     int `json:"totalTask"`
	TotalDoneTask int `json:"totalDoneTask"`
}

func (Goal) FromDomain(goal domain.Goal) Goal {
	return Goal{
		ID:          goal.ID,
		Name:        goal.Name,
		Description: goal.Description,
		Stats: GoalStats{
			TotalTask:     goal.Stats.TotalTask,
			TotalDoneTask: goal.Stats.TotalDoneTask,
		},
		IsCompleted: goal.IsCompleted,
		CreatedAt:   httprespond.NewTimeResponse(goal.CreatedAt),
	}
}
