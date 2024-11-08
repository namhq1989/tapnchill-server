package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

func (s server) registerTaskRoutes() {
	g := s.echo.Group("/api/task")

	g.GET("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetTasksRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetTasks(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetTasksRequest](next)
	})

	g.POST("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateTaskRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateTask(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateTaskRequest](next)
	})

	g.PUT("/:id", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.UpdateTaskRequest)
			performerID = ctx.GetUserID()
			taskID      = c.Param("id")
		)

		resp, err := s.app.UpdateTask(ctx, performerID, taskID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.UpdateTaskRequest](next)
	})

	g.PATCH("/:id/status", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.ChangeTaskStatusRequest)
			performerID = ctx.GetUserID()
			taskID      = c.Param("id")
		)

		resp, err := s.app.ChangeTaskStatus(ctx, performerID, taskID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.ChangeTaskStatusRequest](next)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.DeleteTaskRequest)
			performerID = ctx.GetUserID()
			goalID      = c.Param("id")
		)

		resp, err := s.app.DeleteTask(ctx, performerID, goalID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.DeleteTaskRequest](next)
	})
}
