package sso

import "github.com/namhq1989/go-utilities/appcontext"

type UserInfo struct {
	UID            string                 `json:"uid"`
	Email          string                 `json:"email,omitempty"`
	EmailVerified  bool                   `json:"email_verified,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Picture        string                 `json:"picture,omitempty"`
	ProviderSource string                 `json:"provider_id,omitempty"`
	ProviderUID    string                 `json:"provider_uid,omitempty"`
	Claims         map[string]interface{} `json:"claims,omitempty"`
}

func (s SSO) VerifyToken(ctx *appcontext.AppContext, idToken string) (*UserInfo, error) {
	ctx.Logger().Info("verify Firebase token", appcontext.Fields{"token": idToken})
	token, err := s.firebase.VerifyIDToken(ctx.Context(), idToken)
	if err != nil {
		ctx.Logger().Error("failed to verify Firebase id token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("map token to UserInfo")
	user := UserInfo{
		UID:            token.UID,
		Email:          token.Claims["email"].(string),
		EmailVerified:  token.Claims["email_verified"].(bool),
		Name:           token.Claims["name"].(string),
		Picture:        token.Claims["picture"].(string),
		ProviderSource: token.Claims["firebase"].(map[string]interface{})["sign_in_provider"].(string),
		ProviderUID:    token.Claims["firebase"].(map[string]interface{})["identities"].(map[string]interface{})["google.com"].([]interface{})[0].(string),
		Claims:         token.Claims,
	}

	ctx.Logger().Text("done verify token")
	ctx.Logger().Print("user", user)
	return &user, nil
}
