package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

func (s server) registerCommonRoutes() {
	g := s.echo.Group("/api/common")

	g.GET("/ping", func(c echo.Context) error {
		return httprespond.R200(c, echo.Map{"ok": 1})
	})

	g.POST("/feedback", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateFeedbackRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateFeedback(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateFeedbackRequest](next)
	})

	g.GET("/quote", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetQuoteRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetQuote(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetQuoteRequest](next)
	})

	g.GET("/weather", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetWeatherRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetWeather(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetWeatherRequest](next)
	})
}
