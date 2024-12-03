package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/internal/utils/validation"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

func (s server) registerPaddleRoutes() {
	g := s.echo.Group("/api/webhook")

	g.POST("/paddle", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.PaddleRequest)
		)

		resp, err := s.app.Paddle(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.PaddleRequest](next)
	})

	g.POST("/lemonsqueezy", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.LemonsqueezyRequest)
		)

		resp, err := s.app.Lemonsqueezy(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Read the raw request body
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
			}

			// Retrieve the signature from the headers
			signatureHeader := c.Request().Header.Get("X-Signature")

			// Validate the signature
			if !isValidLemonsqueezySignature(body, s.lemonsqueezySecret, signatureHeader) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid signature"})
			}

			// Reassign the request body for downstream handlers
			c.Request().Body = io.NopCloser(bytes.NewReader(body))

			// Proceed to the next handler
			return next(c)
		}
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.LemonsqueezyRequest](next)
	})
}

func isValidLemonsqueezySignature(body []byte, secret, signatureHeader string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	expectedDigest := h.Sum(nil)
	providedSignature, err := hex.DecodeString(signatureHeader)
	if err != nil {
		return false
	}
	return hmac.Equal(expectedDigest, providedSignature)
}
