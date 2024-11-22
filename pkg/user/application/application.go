package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type (
	Commands interface {
		ExtensionSignIn(ctx *appcontext.AppContext, req dto.ExtensionSignInRequest) (*dto.ExtensionSignInResponse, error)
		GoogleSignIn(ctx *appcontext.AppContext, performerID string, req dto.GoogleSignInRequest) (*dto.GoogleSignInResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.ExtensionSignInHandler
		command.GoogleSignInHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	jwtRepository domain.JwtRepository,
	ssoRepository domain.SSORepository,
	queueRepository domain.QueueRepository,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			ExtensionSignInHandler: command.NewExtensionSignInHandler(userRepository, jwtRepository, queueRepository),
			GoogleSignInHandler:    command.NewGoogleSignInHandler(userRepository, ssoRepository, jwtRepository),
		},
	}
}
