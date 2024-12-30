package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type (
	Commands interface {
		ExtensionSignIn(ctx *appcontext.AppContext, req dto.ExtensionSignInRequest) (*dto.ExtensionSignInResponse, error)
		GoogleSignIn(ctx *appcontext.AppContext, performerID string, req dto.GoogleSignInRequest) (*dto.GoogleSignInResponse, error)
		GenerateSubscriptionCheckoutURL(ctx *appcontext.AppContext, performerID string, req dto.GenerateSubscriptionCheckoutURLRequest) (*dto.GenerateSubscriptionCheckoutURLResponse, error)
	}
	Queries interface {
		GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error)
		GetSubscriptionPlans(ctx *appcontext.AppContext, performerID string, _ dto.GetSubscriptionPlansRequest) (*dto.GetSubscriptionPlansResponse, error)
		GetPaymentCustomerPortalURL(ctx *appcontext.AppContext, performerID string, _ dto.GetPaymentCustomerPortalURLRequest) (*dto.GetPaymentCustomerPortalURLResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.ExtensionSignInHandler
		command.GoogleSignInHandler
		command.GenerateSubscriptionCheckoutURLHandler
	}
	queryHandlers struct {
		query.GetMeHandler
		query.GetSubscriptionPlansHandler
		query.GetPaymentCustomerPortalURLHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	userRepository domain.UserRepository,
	jwtRepository domain.JwtRepository,
	ssoRepository domain.SSORepository,
	queueRepository domain.QueueRepository,
	externalAPIRepository domain.ExternalAPIRepository,
	service domain.Service,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			ExtensionSignInHandler:                 command.NewExtensionSignInHandler(userRepository, jwtRepository, queueRepository),
			GoogleSignInHandler:                    command.NewGoogleSignInHandler(userRepository, ssoRepository, jwtRepository),
			GenerateSubscriptionCheckoutURLHandler: command.NewGenerateSubscriptionCheckoutURLHandler(externalAPIRepository),
		},
		queryHandlers: queryHandlers{
			GetMeHandler:                       query.NewGetMeHandler(service),
			GetSubscriptionPlansHandler:        query.NewGetSubscriptionPlansHandler(),
			GetPaymentCustomerPortalURLHandler: query.NewGetPaymentCustomerPortalURLHandler(service),
		},
	}
}
