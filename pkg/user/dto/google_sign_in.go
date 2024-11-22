package dto

type GoogleSignInRequest struct {
	Token string `json:"token" validate:"required" message:"auth_invalid_google_token"`
}

type GoogleSignInResponse struct {
	UserID      string `json:"userId"`
	Provider    string `json:"provider"`
	AccessToken string `json:"accessToken"`
}
