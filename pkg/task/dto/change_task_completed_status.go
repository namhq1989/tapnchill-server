package dto

type ChangeTaskCompletedStatusRequest struct {
	Completed bool `json:"completed"`
}

type ChangeTaskCompletedStatusResponse struct {
	Completed bool `json:"completed"`
}
