package command

import (
	"fmt"
	"os"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/domain"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

type FastspringHandler struct {
	queueRepository domain.QueueRepository
}

func NewFastspringHandler(queueRepository domain.QueueRepository) FastspringHandler {
	return FastspringHandler{
		queueRepository: queueRepository,
	}
}

func (h FastspringHandler) Fastspring(ctx *appcontext.AppContext, req dto.FastspringRequest) (*dto.FastspringResponse, error) {
	if len(req.Events) == 0 {
		ctx.Logger().Print("invalid payload", req)
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("new FASTSPRING webhook")

	var event = req.Events[0]
	if event.Type == "subscription.activated" {
		return &dto.FastspringResponse{}, h.subscriptionActivated(ctx, event)
	}

	return nil, nil
}

func (h FastspringHandler) subscriptionActivated(ctx *appcontext.AppContext, event dto.FastspringEvent) error {
	ctx.Logger().Info("new FASTSPRING webhook", appcontext.Fields{
		"id": event.ID, "event": event.Type, "account": event.Data.Account,
		"subscription": event.Data.Subscription, "product": event.Data.Product,
		"nextChargeDate": event.Data.NextChargeDate, "tags": event.Data.Tags,
	})

	if event.Data.Tags.UserID == "" {
		ctx.Logger().ErrorText("invalid user id")
		return apperrors.Common.BadRequest
	}

	ctx.Logger().Text("convert nextChargeDate")
	nextBilledAt, err := h.convertNextBilledAt(event.Data.NextChargeDate, event.Data.Product)
	if err != nil {
		ctx.Logger().Error("failed to convent nextChargeDate", err, appcontext.Fields{})
		return apperrors.Common.BadRequest
	}

	items := []string{event.Data.Product}

	ctx.Logger().Text("enqueue subscription activated task")
	if err = h.queueRepository.FastspringSubscriptionActivated(ctx, domain.QueueFastspringSubscriptionActivatedPayload{
		UserID:         event.Data.Tags.UserID,
		SubscriptionID: event.Data.Subscription,
		NextBilledAt:   *nextBilledAt,
		CustomerID:     event.Data.Account,
		Items:          items,
	}); err != nil {
		ctx.Logger().Error("failed to enqueue subscription activated task", err, appcontext.Fields{})
		return err
	}

	return nil
}

func (h FastspringHandler) convertNextBilledAt(nextChargeDate, product string) (*time.Time, error) {
	date, err := time.Parse("2006-01-02", nextChargeDate)
	if err == nil {
		endOfDay := time.Date(
			date.Year(), date.Month(), date.Day(),
			23, 59, 59, 999999999,
			date.Location(),
		)
		return &endOfDay, nil
	}

	// If the date is invalid, calculate based on product type
	now := time.Now()
	switch product {
	case os.Getenv("FASTSPRING_MONTHLY_PRODUCT_ID"):
		oneMonthLater := now.AddDate(0, 1, 0)
		return &oneMonthLater, nil
	case os.Getenv("FASTSPRING_MONTHLY_YEARLY_ID"):
		oneYearLater := now.AddDate(1, 0, 0)
		return &oneYearLater, nil
	default:
		return nil, fmt.Errorf("invalid product type: %s", product)
	}
}
