package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type GetHabitQuotaHandler struct {
	service domain.Service
}

func NewGetHabitQuotaHandler(service domain.Service) GetHabitQuotaHandler {
	return GetHabitQuotaHandler{
		service: service,
	}
}

func (h GetHabitQuotaHandler) GetHabitQuota(ctx *appcontext.AppContext, req *userpb.GetHabitQuotaRequest) (*userpb.GetHabitQuotaResponse, error) {
	ctx.SetTraceID(req.TraceId)
	ctx.Logger().Info("new get habit quota request", appcontext.Fields{"userId": req.UserId})

	limit := domain.FreePlanMaxHabits

	ctx.Logger().Text("find user in db")
	user, err := h.service.GetUserByID(ctx, req.UserId)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return &userpb.GetHabitQuotaResponse{
			Limit:  limit,
			IsFree: true,
		}, nil
	}

	ctx.Logger().Text("check if user is pro plan")
	if user.IsProPlan() {
		ctx.Logger().Text("user is pro plan")
		limit = domain.ProPlanMaxHabits
	}

	ctx.Logger().Text("done get habit quota request")
	return &userpb.GetHabitQuotaResponse{
		Limit:  limit,
		IsFree: user.IsFreePlan(),
	}, nil
}
