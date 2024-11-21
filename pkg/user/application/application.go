package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type (
	Commands interface {
		AnonymousSignIn(ctx *appcontext.AppContext, req dto.AnonymousSignInRequest) (*dto.AnonymousSignInResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.AnonymousSignInHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	jwtRepository domain.JwtRepository,
	queueRepository domain.QueueRepository,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			AnonymousSignInHandler: command.NewAnonymousSignInHandler(userRepository, jwtRepository, queueRepository),
		},
	}
}
