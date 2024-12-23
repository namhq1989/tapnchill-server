package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type GetQRCodeQuotaHandler struct {
	service domain.Service
}

func NewGetQRCodeQuotaHandler(service domain.Service) GetQRCodeQuotaHandler {
	return GetQRCodeQuotaHandler{
		service: service,
	}
}

func (h GetQRCodeQuotaHandler) GetQRCodeQuota(ctx *appcontext.AppContext, req *userpb.GetQRCodeQuotaRequest) (*userpb.GetQRCodeQuotaResponse, error) {
	ctx.SetTraceID(req.TraceId)
	ctx.Logger().Info("new get QR code quota request", appcontext.Fields{"userId": req.UserId})

	limit := domain.FreePlanMaxQRCodes

	ctx.Logger().Text("find user in db")
	user, err := h.service.GetUserByID(ctx, req.UserId)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return &userpb.GetQRCodeQuotaResponse{
			Limit:  limit,
			IsFree: true,
		}, nil
	}

	ctx.Logger().Text("check if user is pro plan")
	if user.IsProPlan() {
		ctx.Logger().Text("user is pro plan")
		limit = domain.ProPlanMaxQRCodes
	}

	ctx.Logger().Text("done get QR code quota request")
	return &userpb.GetQRCodeQuotaResponse{
		Limit:  limit,
		IsFree: user.IsFreePlan(),
	}, nil
}
