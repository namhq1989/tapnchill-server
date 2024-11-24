package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

type (
	Commands interface {
		Paddle(ctx *appcontext.AppContext, req dto.PaddleRequest) (*dto.PaddleResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.PaddleHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New() *Application {
	return &Application{
		commandHandlers: commandHandlers{
			PaddleHandler: command.NewPaddleHandler(),
		},
	}
}
