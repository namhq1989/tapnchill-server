package user

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/user/application"
	"github.com/namhq1989/tapnchill-server/pkg/user/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/user/rest"
	"github.com/namhq1989/tapnchill-server/pkg/user/shared"
	"github.com/namhq1989/tapnchill-server/pkg/user/worker"
)

type Module struct{}

func (Module) Name() string {
	return "USER"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		userRepository                = infrastructure.NewUserRepository(mono.Database(), mono.Config().AnonymousUserChecksumSecret)
		subscriptionHistoryRepository = infrastructure.NewSubscriptionHistoryRepository(mono.Database())

		jwtRepository     = infrastructure.NewJwtRepository(mono.JWT())
		ssoRepository     = infrastructure.NewSSORepository(mono.SSO())
		queueRepository   = infrastructure.NewQueueRepository(mono.Queue())
		cachingRepository = infrastructure.NewCachingRepository(mono.Caching())

		service = shared.NewService(userRepository, cachingRepository)

		// app
		app = application.New(
			userRepository,
			jwtRepository,
			ssoRepository,
			queueRepository,
			service,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	w := worker.New(
		mono.Queue(),
		userRepository,
		subscriptionHistoryRepository,
		cachingRepository,
	)
	w.Start()

	return nil
}
