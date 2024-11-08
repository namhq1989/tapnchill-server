package dto

import "time"

type CreateTaskRequest struct {
	GoalID      string     `json:"goalId" validate:"required" message:"task_invalid_goal_id"`
	Name        string     `json:"name" validate:"required" message:"invalid_name"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}

type CreateTaskResponse struct {
	ID string `json:"id"`
}
