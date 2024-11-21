package domain

import "github.com/namhq1989/go-utilities/appcontext"

type SSORepository interface {
	VerifyGoogleToken(ctx *appcontext.AppContext, token string) (*SSOGoogleUser, error)
}

type SSOGoogleUser struct {
	UID   string
	Email string
	Name  string
}
