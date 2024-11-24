package habit

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/grpcclient"
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
	userGRPCClient, err := grpcclient.NewUserClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		// dependencies
		habitRepository           = infrastructure.NewHabitRepository(mono.Database())
		habitCompletionRepository = infrastructure.NewHabitCompletionRepository(mono.Database())
		habitDailyStatsRepository = infrastructure.NewHabitDailyStatsRepository(mono.Database())
		cachingRepository         = infrastructure.NewCachingRepository(mono.Caching())

		service = shared.NewService(habitRepository, habitDailyStatsRepository, cachingRepository)

		userHub = infrastructure.NewUserHub(userGRPCClient)

		// app
		app = application.New(
			habitRepository,
			habitCompletionRepository,
			habitDailyStatsRepository,
			service,
			userHub,
		)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
