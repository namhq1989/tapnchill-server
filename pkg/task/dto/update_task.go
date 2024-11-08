package dto

import "time"

type UpdateTaskRequest struct {
	Name        string     `json:"name" validate:"required" message:"invalid_name"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}

type UpdateTaskResponse struct {
	ID string `json:"id"`
}
