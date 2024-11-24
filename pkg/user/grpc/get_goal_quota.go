package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type GetGoalQuotaHandler struct {
	service domain.Service
}

func NewGetGoalQuotaHandler(service domain.Service) GetGoalQuotaHandler {
	return GetGoalQuotaHandler{
		service: service,
	}
}

func (h GetGoalQuotaHandler) GetGoalQuota(ctx *appcontext.AppContext, req *userpb.GetGoalQuotaRequest) (*userpb.GetGoalQuotaResponse, error) {
	ctx.SetTraceID(req.TraceId)
	ctx.Logger().Info("new get goal quota request", appcontext.Fields{"userId": req.UserId})

	limit := domain.FreePlanMaxGoals

	ctx.Logger().Text("find user in db")
	user, err := h.service.GetUserByID(ctx, req.UserId)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return &userpb.GetGoalQuotaResponse{
			Limit: limit,
		}, nil
	}

	ctx.Logger().Text("check if user is pro plan")
	if user.IsProPlan() {
		ctx.Logger().Text("user is pro plan")
		limit = domain.ProPlanMaxGoals
	}

	ctx.Logger().Text("done get goal quota request")
	return &userpb.GetGoalQuotaResponse{
		Limit: limit,
	}, nil
}
