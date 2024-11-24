package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type GetTaskQuotaHandler struct {
	service domain.Service
}

func NewGetTaskQuotaHandler(service domain.Service) GetTaskQuotaHandler {
	return GetTaskQuotaHandler{
		service: service,
	}
}

func (h GetTaskQuotaHandler) GetTaskQuota(ctx *appcontext.AppContext, req *userpb.GetTaskQuotaRequest) (*userpb.GetTaskQuotaResponse, error) {
	ctx.SetTraceID(req.TraceId)
	ctx.Logger().Info("new get task quota request", appcontext.Fields{"userId": req.UserId})

	limit := domain.FreePlanMaxTaskPerGoal

	ctx.Logger().Text("find user in db")
	user, err := h.service.GetUserByID(ctx, req.UserId)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return &userpb.GetTaskQuotaResponse{
			Limit: limit,
		}, nil
	}

	ctx.Logger().Text("check if user is pro plan")
	if user.IsProPlan() {
		ctx.Logger().Text("user is pro plan")
		limit = domain.ProPlanMaxTaskPerGoal
	}

	ctx.Logger().Text("done get task quota request")
	return &userpb.GetTaskQuotaResponse{
		Limit: limit,
	}, nil
}
