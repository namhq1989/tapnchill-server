package note

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/grpcclient"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/note/application"
	"github.com/namhq1989/tapnchill-server/pkg/note/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/note/rest"
)

type Module struct{}

func (Module) Name() string {
	return "NOTE"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	userGRPCClient, err := grpcclient.NewUserClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		// dependencies
		noteRepository = infrastructure.NewNoteRepository(mono.Database())

		userHub = infrastructure.NewUserHub(userGRPCClient)

		// app
		app = application.New(
			noteRepository,
			userHub,
		)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
