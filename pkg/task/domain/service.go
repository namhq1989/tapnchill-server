package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetTaskByID(ctx *appcontext.AppContext, taskID, userID string) (*Task, error)
	GetGoalByID(ctx *appcontext.AppContext, goalID, userID string) (*Goal, error)
}
