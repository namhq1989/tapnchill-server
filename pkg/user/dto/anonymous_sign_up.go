package dto

type AnonymousSignUpRequest struct {
	ClientID string `json:"clientId" validate:"required" message:"user_invalid_client_id"`
	Source   string `json:"source"`
	Checksum string `json:"checksum"`
}

type AnonymousSignUpResponse struct {
	AccessToken string `json:"accessToken"`
}
