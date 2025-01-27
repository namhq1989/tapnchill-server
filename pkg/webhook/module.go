package webhook

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/application"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/rest"
)

type Module struct{}

func (Module) Name() string {
	return "WEBHOOK"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		queueRepository = infrastructure.NewQueueRepository(mono.Queue())

		// app
		app = application.New(queueRepository)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().LemonsqueezySigningSecret, mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
