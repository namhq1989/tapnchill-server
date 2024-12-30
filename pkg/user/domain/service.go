package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetUserByID(ctx *appcontext.AppContext, userID string) (*User, error)
	GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, userID string) (*string, error)
}
