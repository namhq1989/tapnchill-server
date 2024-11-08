package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetGoalByID(ctx *appcontext.AppContext, goalID, userID string) (*Goal, error)
	GetTaskByID(ctx *appcontext.AppContext, taskID, userID string) (*Task, error)
}
