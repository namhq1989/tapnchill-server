package main

import (
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/config"
	"golang.org/x/text/language"
)

func initRest(cfg config.Server) *echo.Echo {
	// echo instance
	e := echo.New()

	setMiddleware(e, cfg)
	return e
}

func setMiddleware(e *echo.Echo, cfg config.Server) {
	addCorsMiddleware(e, cfg)
	addContext(e)
	addIp(e)
	addLanguageMiddleware(e)
	addRateLimiter(e)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.Secure())

	if cfg.IsEnvRelease {
		e.Use(middleware.Recover())
	} else {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
		}))
	}
}

func addContext(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := appcontext.NewRest(c.Request().Context())
			c.Set("ctx", ctx)

			return next(c)
		}
	})
}

var localIps = []string{
	"127.0.0.1",
	"::ffff:127.0.0.1",
	"::1",
}

func addIp(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				ctx = c.Get("ctx").(*appcontext.AppContext)
				ip  = c.RealIP()
			)

			if slices.Contains(localIps, ip) {
				ip = "171.225.184.76" // set default to DN for local testing
			}

			ctx.SetIP(ip)
			return next(c)
		}
	})
}

func addLanguageMiddleware(e *echo.Echo) {
	supportedLanguages := language.NewMatcher([]language.Tag{
		language.English,
		language.Vietnamese,
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Get("ctx").(*appcontext.AppContext)

			// parse the Accept-Language header
			accept := c.Request().Header.Get("Accept-Language")
			tag, _, _ := language.ParseAcceptLanguage(accept)

			// match the best supported language
			matched, _, _ := supportedLanguages.Match(tag...)

			// Use "en" as default if no match
			lang := language.Vietnamese.String()
			if matched == language.English {
				lang = language.English.String()
			}

			// set the language in the context
			ctx.SetLang(lang)

			// Call the next handler in the chain
			return next(c)
		}
	})
}

func addCorsMiddleware(e *echo.Echo, cfg config.Server) {
	allowedOrigins := []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"chrome-extension://inncdblgnijbdpedbfcoieknfgpkflnd",
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func addRateLimiter(e *echo.Echo) {
	cfg := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 60, Burst: 60, ExpiresIn: 5 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(cfg))
}