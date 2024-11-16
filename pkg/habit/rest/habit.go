package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

func (s server) registerHabitRoutes() {
	g := s.echo.Group("/api/habit")

	g.GET("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetHabitsRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetHabits(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetHabitsRequest](next)
	})

	g.POST("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateHabitRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateHabit(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateHabitRequest](next)
	})

	g.PUT("/:id", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.UpdateHabitRequest)
			performerID = ctx.GetUserID()
			habitID     = c.Param("id")
		)

		resp, err := s.app.UpdateHabit(ctx, performerID, habitID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.UpdateHabitRequest](next)
	})

	g.PATCH("/:id/status", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.ChangeHabitStatusRequest)
			performerID = ctx.GetUserID()
			habitID     = c.Param("id")
		)

		resp, err := s.app.ChangeHabitStatus(ctx, performerID, habitID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.ChangeHabitStatusRequest](next)
	})

	g.POST("/:id/complete", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CompleteHabitRequest)
			performerID = ctx.GetUserID()
			habitID     = c.Param("id")
		)

		resp, err := s.app.CompleteHabit(ctx, performerID, habitID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CompleteHabitRequest](next)
	})

	g.GET("/stat", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetStatsRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetStats(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetStatsRequest](next)
	})
}
