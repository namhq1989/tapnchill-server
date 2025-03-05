package domain

import "github.com/namhq1989/go-utilities/appcontext"

type ReportRepository interface {
	NewUserSignedInWithGoogle(ctx *appcontext.AppContext, userID, email string) error
}
