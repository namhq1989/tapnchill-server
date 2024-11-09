package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

func (s server) registerUserRoutes() {
	g := s.echo.Group("/api/user")

	g.POST("/sign-up/anonymous", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.AnonymousSignUpRequest)
		)

		resp, err := s.app.AnonymousSignUp(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.AnonymousSignUpRequest](next)
	})
}
