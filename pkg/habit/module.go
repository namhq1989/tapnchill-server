package habit

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application"
	"github.com/namhq1989/tapnchill-server/pkg/habit/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/habit/rest"
	"github.com/namhq1989/tapnchill-server/pkg/habit/shared"
)

type Module struct{}

func (Module) Name() string {
	return "HABIT"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		habitRepository           = infrastructure.NewHabitRepository(mono.Database())
		habitCompletionRepository = infrastructure.NewHabitCompletionRepository(mono.Database())
		habitDailyStatsRepository = infrastructure.NewHabitDailyStatsRepository(mono.Database())

		service = shared.NewService(habitRepository)

		// app
		app = application.New(
			habitRepository,
			habitCompletionRepository,
			habitDailyStatsRepository,
			service,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
