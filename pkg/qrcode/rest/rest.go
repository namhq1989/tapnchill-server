package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	appjwt "github.com/namhq1989/tapnchill-server/internal/jwt"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application"
)

type server struct {
	app          application.Instance
	echo         *echo.Echo
	jwt          appjwt.Operations
	isEnvRelease bool
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt *appjwt.JWT, isEnvRelease bool) error {
	var s = server{
		app:          app,
		echo:         e,
		jwt:          jwt,
		isEnvRelease: isEnvRelease,
	}

	s.registerQRCodeRoutes()

	return nil
}
