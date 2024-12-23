package domain

import "github.com/namhq1989/go-utilities/appcontext"

type UserHub interface {
	GetGoalQuota(ctx *appcontext.AppContext, userID string) (int64, bool, error)
	GetTaskQuota(ctx *appcontext.AppContext, userID string) (int64, bool, error)
}
