package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type (
	Commands interface {
		CreateGoal(ctx *appcontext.AppContext, performerID string, req dto.CreateGoalRequest) (*dto.CreateGoalResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.CreateGoalHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	goalRepository domain.GoalRepository,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateGoalHandler: command.NewCreateGoalHandler(goalRepository),
		},
	}
}
