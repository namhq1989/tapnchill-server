package dto

type GoogleSignInRequest struct {
	Token string `json:"token" validate:"required" message:"auth_invalid_google_token"`
}

type GoogleSignInResponse struct {
	AccessToken string `json:"accessToken"`
}
