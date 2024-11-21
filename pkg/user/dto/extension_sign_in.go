package dto

type ExtensionSignInRequest struct {
	ClientID string `json:"clientId" validate:"required" message:"user_invalid_client_id"`
	Checksum string `json:"checksum"`
}

type ExtensionSignInResponse struct {
	AccessToken string `json:"accessToken"`
}
