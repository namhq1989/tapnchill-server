package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

func (s server) registerGoalRoutes() {
	g := s.echo.Group("/api/task")

	g.POST("/goal", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateGoalRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateGoal(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateGoalRequest](next)
	})
}
