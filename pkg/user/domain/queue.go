package domain

import "github.com/namhq1989/go-utilities/appcontext"

type QueueRepository interface {
	CreateUserDefaultGoal(ctx *appcontext.AppContext, payload QueueCreateUserDefaultGoalPayload) error
}

type QueueCreateUserDefaultGoalPayload struct {
	UserID string
}
