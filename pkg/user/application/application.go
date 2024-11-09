package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type (
	Commands interface {
		AnonymousSignUp(ctx *appcontext.AppContext, req dto.AnonymousSignUpRequest) (*dto.AnonymousSignUpResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.AnonymousSignUpHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	jwtRepository domain.JwtRepository,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			AnonymousSignUpHandler: command.NewAnonymousSignUpHandler(userRepository, jwtRepository),
		},
	}
}
