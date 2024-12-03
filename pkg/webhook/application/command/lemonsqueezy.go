package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/domain"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

type LemonsqueezyHandler struct {
	queueRepository domain.QueueRepository
}

func NewLemonsqueezyHandler(queueRepository domain.QueueRepository) LemonsqueezyHandler {
	return LemonsqueezyHandler{
		queueRepository: queueRepository,
	}
}

func (h LemonsqueezyHandler) Lemonsqueezy(ctx *appcontext.AppContext, req dto.LemonsqueezyRequest) (*dto.LemonsqueezyResponse, error) {
	ctx.Logger().Text("new LEMONSQUEEZY webhook")

	if req.Meta.EventName == "subscription_payment_success" {
		return &dto.LemonsqueezyResponse{}, h.subscriptionPaymentSuccess(ctx, req)
	}

	return nil, nil
}

func (h LemonsqueezyHandler) subscriptionPaymentSuccess(ctx *appcontext.AppContext, req dto.LemonsqueezyRequest) error {
	ctx.Logger().Info("new subscription payment success", appcontext.Fields{
		"invoiceID": req.Data.InvoiceID, "userID": req.Meta.CustomData.UserID,
	})

	ctx.Logger().Text("enqueue subscription payment success task")
	if err := h.queueRepository.LemonsqueezySubscriptionPaymentSuccess(ctx, domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload{
		UserID:    req.Meta.CustomData.UserID,
		InvoiceID: req.Data.InvoiceID,
	}); err != nil {
		ctx.Logger().Error("failed to enqueue subscription payment success task", err, appcontext.Fields{})
		return err
	}

	return nil
}
