package dto

type UpdateGoalRequest struct {
	Name        string `json:"name" validate:"required" message:"invalid_name"`
	Description string `json:"description"`
}

type UpdateGoalResponse struct {
	ID string `json:"id"`
}
