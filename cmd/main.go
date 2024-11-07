package main

import (
	"crypto/subtle"
	"fmt"
	"time"

	"github.com/namhq1989/tapnchill-server/internal/externalapi"

	"github.com/namhq1989/tapnchill-server/pkg/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/go-utilities/logger"
	"github.com/namhq1989/tapnchill-server/internal/caching"
	"github.com/namhq1989/tapnchill-server/internal/config"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	appjwt "github.com/namhq1989/tapnchill-server/internal/jwt"
	"github.com/namhq1989/tapnchill-server/internal/monolith"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/internal/utils/waiter"
)

func main() {
	var err error

	// config
	cfg := config.Init()

	// logger
	logger.Init(cfg.Environment)

	// app error
	apperrors.Init()

	// server
	a := app{}
	a.cfg = cfg

	// jwt
	a.jwt, err = appjwt.Init(cfg.AccessTokenSecret, time.Second*time.Duration(cfg.AccessTokenTTL))
	if err != nil {
		panic(err)
	}

	// rest
	a.rest = initRest(cfg)

	// grpc
	a.rpc = initRPC()

	// database
	a.database = database.NewDatabaseClient(cfg.MongoURL, cfg.MongoDBName)

	// caching
	a.caching = caching.NewCachingClient(cfg.CachingRedisURL)

	// external api
	a.externalApi = externalapi.NewExternalAPIClient()

	// queue
	a.queue = queue.Init(cfg.QueueRedisURL, cfg.QueueConcurrency)

	// init queue's dashboard
	a.rest.Any(fmt.Sprintf("%s/*", queue.DashboardPath), echo.WrapHandler(queue.EnableDashboard(cfg.QueueRedisURL)), middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
		if !cfg.IsEnvRelease {
			return true, nil
		}
		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// waiter
	a.waiter = waiter.New(waiter.CatchSignals())

	// modules
	a.modules = []monolith.Module{
		&common.Module{},
	}

	// start
	if err = a.startupModules(); err != nil {
		panic(err)
	}

	fmt.Println("--- started tapnchill-server application")
	defer fmt.Println("--- stopped tapnchill-server application")

	// wait for other service starts
	a.waiter.Add(
		a.waitForRest,
		a.waitForRPC,
	)
	if err = a.waiter.Wait(); err != nil {
		panic(err)
	}
}
