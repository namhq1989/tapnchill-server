package dto

type ChangeHabitStatusRequest struct {
	Date   string `json:"date" validate:"required" message:"invalid_date"`
	Status string `json:"status" validate:"required" message:"invalid_status"`
}

type ChangeHabitStatusResponse struct{}
