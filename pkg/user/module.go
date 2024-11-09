package user

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/user/application"
	"github.com/namhq1989/tapnchill-server/pkg/user/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/user/rest"
)

type Module struct{}

func (Module) Name() string {
	return "USER"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		userRepository = infrastructure.NewUserRepository(mono.Database(), mono.Config().AnonymousUserChecksumSecret)

		jwtRepository = infrastructure.NewJwtRepository(mono.JWT())

		// app
		app = application.New(
			userRepository,
			jwtRepository,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
