package domain

import "github.com/namhq1989/go-utilities/appcontext"

type CachingRepository interface {
	GetUserByID(ctx *appcontext.AppContext, userID string) (*User, error)
	SetUserByID(ctx *appcontext.AppContext, userID string, user User) error
	DeleteUserByID(ctx *appcontext.AppContext, userID string) error

	GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, customerID string) (*string, error)
	SetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, customerID, url string) error
}
