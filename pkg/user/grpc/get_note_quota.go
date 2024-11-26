package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type GetNoteQuotaHandler struct {
	service domain.Service
}

func NewGetNoteQuotaHandler(service domain.Service) GetNoteQuotaHandler {
	return GetNoteQuotaHandler{
		service: service,
	}
}

func (h GetNoteQuotaHandler) GetNoteQuota(ctx *appcontext.AppContext, req *userpb.GetNoteQuotaRequest) (*userpb.GetNoteQuotaResponse, error) {
	ctx.SetTraceID(req.TraceId)
	ctx.Logger().Info("new get note quota request", appcontext.Fields{"userId": req.UserId})

	limit := domain.FreePlanMaxNotes

	ctx.Logger().Text("find user in db")
	user, err := h.service.GetUserByID(ctx, req.UserId)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return &userpb.GetNoteQuotaResponse{
			Limit: limit,
		}, nil
	}

	ctx.Logger().Text("check if user is pro plan")
	if user.IsProPlan() {
		ctx.Logger().Text("user is pro plan")
		limit = domain.ProPlanMaxNotes
	}

	ctx.Logger().Text("done get note quota request")
	return &userpb.GetNoteQuotaResponse{
		Limit: limit,
	}, nil
}
