package common

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/common/application"
	"github.com/namhq1989/tapnchill-server/pkg/common/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/common/rest"
	"github.com/namhq1989/tapnchill-server/pkg/common/worker"
)

type Module struct{}

func (Module) Name() string {
	return "COMMON"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		feedbackRepository = infrastructure.NewFeedbackRepository(mono.Database())
		quoteRepository    = infrastructure.NewQuoteRepository(mono.Database())

		externalApiRepository = infrastructure.NewExternalAPIRepository(mono.ExternalApi())

		// app
		app = application.New(
			feedbackRepository,
			quoteRepository,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	// worker
	w := worker.New(
		mono.Queue(),
		quoteRepository,
		externalApiRepository,
	)
	w.Start()

	return nil
}
