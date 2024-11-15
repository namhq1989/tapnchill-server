package dto

type UpdateHabitRequest struct {
	Name       string `json:"name" validate:"required" message:"invalid_name"`
	Goal       string `json:"goal" validate:"required" message:"invalid_goal"`
	DayOfWeeks []int  `json:"dayOfWeeks" validate:"required" message:"invalid_days_of_week"`
	Icon       string `json:"icon" validate:"required" message:"invalid_icon"`
}

type UpdateHabitResponse struct {
	ID string `json:"id"`
}
