package dto

type CompleteHabitRequest struct {
	Date string `json:"date" validate:"required" message:"invalid_date"`
}

type CompleteHabitResponse struct {
	ID string `json:"id"`
}
