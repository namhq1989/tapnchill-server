package task

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/task/application"
	"github.com/namhq1989/tapnchill-server/pkg/task/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/task/rest"
	"github.com/namhq1989/tapnchill-server/pkg/task/shared"
	"github.com/namhq1989/tapnchill-server/pkg/task/worker"
)

type Module struct{}

func (Module) Name() string {
	return "TASK"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// dependencies
		taskRepository = infrastructure.NewTaskRepository(mono.Database())
		goalRepository = infrastructure.NewGoalRepository(mono.Database())

		service = shared.NewService(taskRepository, goalRepository)

		// app
		app = application.New(
			taskRepository,
			goalRepository,
			service,
		)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	w := worker.New(
		mono.Queue(),
		goalRepository,
	)
	w.Start()

	return nil
}
