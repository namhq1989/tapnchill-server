package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	appjwt "github.com/namhq1989/tapnchill-server/internal/jwt"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/application"
)

type server struct {
	app                application.Instance
	echo               *echo.Echo
	jwt                appjwt.Operations
	lemonsqueezySecret string
	isEnvRelease       bool
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt *appjwt.JWT, lemonsqueezySecret string, isEnvRelease bool) error {
	var s = server{
		app:                app,
		echo:               e,
		jwt:                jwt,
		lemonsqueezySecret: lemonsqueezySecret,
		isEnvRelease:       isEnvRelease,
	}

	s.registerPaddleRoutes()

	return nil
}
