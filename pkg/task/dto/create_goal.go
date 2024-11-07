package dto

type CreateGoalRequest struct {
	Name        string `json:"name" validate:"required" message:"invalid_name"`
	Description string `json:"description" validate:"required" message:"invalid_description"`
}

type CreateGoalResponse struct {
	ID string `json:"id"`
}
