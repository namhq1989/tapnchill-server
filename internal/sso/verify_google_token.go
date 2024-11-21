package sso

import "github.com/namhq1989/go-utilities/appcontext"

type UserInfo struct {
	UID    string                 `json:"uid"`
	Email  string                 `json:"email,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Claims map[string]interface{} `json:"claims,omitempty"`
}

func (s SSO) VerifyGoogleToken(ctx *appcontext.AppContext, idToken string) (*UserInfo, error) {
	ctx.Logger().Info("verify Firebase token", appcontext.Fields{"token": idToken})
	token, err := s.firebase.VerifyIDToken(ctx.Context(), idToken)
	if err != nil {
		ctx.Logger().Error("failed to verify Firebase id token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("map token to UserInfo")
	user := UserInfo{
		UID:    token.UID,
		Email:  token.Claims["email"].(string),
		Name:   token.Claims["name"].(string),
		Claims: token.Claims,
	}

	ctx.Logger().Text("done verify token")
	ctx.Logger().Print("user", user)
	return &user, nil
}
