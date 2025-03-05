package domain

import "github.com/namhq1989/go-utilities/appcontext"

type ReportRepository interface {
	NewUserFeedback(ctx *appcontext.AppContext, feedbackID, content string) error
}
