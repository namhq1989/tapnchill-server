package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

func (s server) registerQRCodeRoutes() {
	g := s.echo.Group("/api/qr-code")

	g.GET("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetQRCodesRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetQRCodes(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetQRCodesRequest](next)
	})

	g.POST("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.CreateQRCodeRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.CreateQRCode(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.CreateQRCodeRequest](next)
	})

	g.PUT("/:id", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.UpdateQRCodeRequest)
			performerID = ctx.GetUserID()
			qrCodeID    = c.Param("id")
		)

		resp, err := s.app.UpdateQRCode(ctx, performerID, qrCodeID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.UpdateQRCodeRequest](next)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.DeleteQRCodeRequest)
			performerID = ctx.GetUserID()
			qrCodeID    = c.Param("id")
		)

		resp, err := s.app.DeleteQRCode(ctx, performerID, qrCodeID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.DeleteQRCodeRequest](next)
	})
}
