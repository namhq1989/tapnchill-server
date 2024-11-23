package domain

import "github.com/namhq1989/go-utilities/appcontext"

type CachingRepository interface {
	GetUserByID(ctx *appcontext.AppContext, userID string) (*User, error)
	SetUserByID(ctx *appcontext.AppContext, userID string, user User) error
	DeleteUserByID(ctx *appcontext.AppContext, userID string) error
}
