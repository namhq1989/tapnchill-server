package dto

type UpdateQRCodeRequest struct {
	Name string `json:"name" validate:"required" message:"invalid_name"`
}

type UpdateQRCodeResponse struct {
	ID string `json:"id"`
}
