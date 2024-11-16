package dto

type ChangeHabitStatusRequest struct {
	Status string `json:"status" validate:"required" message:"invalid_status"`
}

type ChangeHabitStatusResponse struct{}
