package domain

import "github.com/namhq1989/go-utilities/appcontext"

type SSORepository interface {
	GetUserInfoWithToken(ctx *appcontext.AppContext, token string) (*SSOUser, error)
}

type SSOUser struct {
	UID            string
	Email          string
	Name           string
	ProviderSource string
	ProviderUID    string
}
