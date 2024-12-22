package qrcode

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/grpcclient"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/infrastructure"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/rest"
)

type Module struct{}

func (Module) Name() string {
	return "QR CODE"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	userGRPCClient, err := grpcclient.NewUserClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		// dependencies
		qrCodeRepository = infrastructure.NewQRCodeRepository(mono.Database())

		userHub = infrastructure.NewUserHub(userGRPCClient)

		// app
		app = application.New(
			qrCodeRepository,
			userHub,
		)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), mono.Config().IsEnvRelease); err != nil {
		return err
	}

	return nil
}
