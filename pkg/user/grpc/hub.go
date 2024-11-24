package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type (
	Hubs interface {
		GetHabitQuota(ctx *appcontext.AppContext, req *userpb.GetHabitQuotaRequest) (*userpb.GetHabitQuotaResponse, error)
		GetGoalQuota(ctx *appcontext.AppContext, req *userpb.GetGoalQuotaRequest) (*userpb.GetGoalQuotaResponse, error)
		GetTaskQuota(ctx *appcontext.AppContext, req *userpb.GetTaskQuotaRequest) (*userpb.GetTaskQuotaResponse, error)
	}
	App interface {
		Hubs
	}

	appHubHandler struct {
		GetHabitQuotaHandler
		GetGoalQuotaHandler
		GetTaskQuotaHandler
	}
	Application struct {
		appHubHandler
	}
)

var _ App = (*Application)(nil)

func New(
	service domain.Service,
) *Application {
	return &Application{
		appHubHandler: appHubHandler{
			GetHabitQuotaHandler: NewGetHabitQuotaHandler(service),
			GetGoalQuotaHandler:  NewGetGoalQuotaHandler(service),
			GetTaskQuotaHandler:  NewGetTaskQuotaHandler(service),
		},
	}
}
