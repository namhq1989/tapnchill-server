package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type PaddleSubscriptionCreatedHandler struct {
	userRepository                domain.UserRepository
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository
}

func NewPaddleSubscriptionCreatedHandler(
	userRepository domain.UserRepository,
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository,
) PaddleSubscriptionCreatedHandler {
	return PaddleSubscriptionCreatedHandler{
		userRepository:                userRepository,
		subscriptionHistoryRepository: subscriptionHistoryRepository,
	}
}

func (h PaddleSubscriptionCreatedHandler) PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueuePaddleSubscriptionCreatedPayload) error {
	ctx.Logger().Text("find user in db")
	user, err := h.userRepository.FindByID(ctx, payload.UserID)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return nil
	}

	ctx.Logger().Text("create new subscription history model")
	history, err := domain.NewSubscriptionHistory(payload.UserID, payload.SubscriptionID, domain.PaymentSourcePaddle, payload.CustomerID, payload.Items, payload.NextBilledAt)
	if err != nil {
		ctx.Logger().Error("failed to create new subscription history model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist subscription history in db")
	if err = h.subscriptionHistoryRepository.Create(ctx, *history); err != nil {
		ctx.Logger().Error("failed to persist subscription history in db", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update user subscription customer id")
	user.SetSubscriptionCustomerID(payload.CustomerID)

	ctx.Logger().Text("update user in db")
	if err = h.userRepository.Update(ctx, *user); err != nil {
		ctx.Logger().Error("failed to update user in db", err, appcontext.Fields{})
		return err
	}

	return nil
}
