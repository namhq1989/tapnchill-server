package dto

type AnonymousSignUpRequest struct {
	ClientID string `json:"clientId" validate:"required" message:"user_invalid_client_id"`
	Checksum string `json:"checksum"`
}

type AnonymousSignUpResponse struct {
	AccessToken string `json:"accessToken"`
}
