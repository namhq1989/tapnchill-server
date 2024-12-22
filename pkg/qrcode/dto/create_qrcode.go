package dto

type CreateQRCodeRequest struct {
	Name     string         `json:"name" validate:"required" message:"invalid_name"`
	Type     string         `json:"type" validate:"required" message:"invalid_type"`
	Content  string         `json:"content" validate:"required" message:"invalid_content"`
	Settings QRCodeSettings `json:"settings" validate:"required" message:"invalid_qr_code_settings"`
	Data     string         `json:"data"`
}

type CreateQRCodeResponse struct {
	ID string `json:"id"`
}
