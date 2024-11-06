package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/caching"
	"github.com/namhq1989/tapnchill-server/internal/config"
	"github.com/namhq1989/tapnchill-server/internal/database"
	appjwt "github.com/namhq1989/tapnchill-server/internal/jwt"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Caching() *caching.Caching
	JWT() *appjwt.JWT
	Queue() *queue.Queue
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
