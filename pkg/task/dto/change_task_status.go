package dto

type ChangeTaskStatusRequest struct {
	Status string `json:"status"`
}

type ChangeTaskStatusResponse struct{}
