package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/domain"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

type PaddleHandler struct {
	queueRepository domain.QueueRepository
}

func NewPaddleHandler(queueRepository domain.QueueRepository) PaddleHandler {
	return PaddleHandler{
		queueRepository: queueRepository,
	}
}

func (h PaddleHandler) Paddle(ctx *appcontext.AppContext, req dto.PaddleRequest) (*dto.PaddleResponse, error) {
	ctx.Logger().Text("new PADDLE webhook")

	if req.EventType == "subscription.created" {
		return &dto.PaddleResponse{}, h.subscriptionCreated(ctx, req)
	} else if req.EventType == "transaction.completed" {
		return &dto.PaddleResponse{}, h.transactionCompleted(ctx, req)
	}

	return nil, nil
}

func (h PaddleHandler) subscriptionCreated(ctx *appcontext.AppContext, req dto.PaddleRequest) error {
	ctx.Logger().Info("new subscription created", appcontext.Fields{
		"id": req.Data.ID, "subscriptionID": req.Data.SubscriptionID, "customerID": req.Data.CustomerID,
		"nextBilledAt": req.Data.NextBilledAt, "items": req.Data.Items, "customData": req.Data.CustomData,
	})

	nextBilledAt, err := manipulation.GetEndOfDayWithClientDate(req.Data.NextBilledAt)
	if err != nil {
		ctx.Logger().Error("failed to get end of day with client date", err, appcontext.Fields{})
		return apperrors.Common.BadRequest
	}

	items := make([]string, len(req.Data.Items))
	for i, item := range req.Data.Items {
		items[i] = item.Price.ID
	}

	ctx.Logger().Text("enqueue subscription created task")
	if err = h.queueRepository.PaddleSubscriptionCreated(ctx, domain.QueuePaddleSubscriptionCreatedPayload{
		UserID:         req.Data.CustomData.UserID,
		SubscriptionID: req.Data.ID,
		NextBilledAt:   *nextBilledAt,
		CustomerID:     req.Data.CustomerID,
		Items:          items,
	}); err != nil {
		ctx.Logger().Error("failed to enqueue subscription created task", err, appcontext.Fields{})
		return err
	}

	return nil
}

func (h PaddleHandler) transactionCompleted(ctx *appcontext.AppContext, req dto.PaddleRequest) error {
	ctx.Logger().Info("new transaction completed", appcontext.Fields{
		"id": req.Data.ID, "subscriptionID": req.Data.SubscriptionID,
	})

	ctx.Logger().Text("enqueue transaction completed task")
	if err := h.queueRepository.PaddleTransactionCompleted(ctx, domain.QueuePaddleTransactionCompletedPayload{
		UserID:         req.Data.CustomData.UserID,
		SubscriptionID: req.Data.SubscriptionID,
	}); err != nil {
		ctx.Logger().Error("failed to enqueue transaction completed task", err, appcontext.Fields{})
		return err
	}

	return nil
}
