package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"go.opentelemetry.io/otel"
)

type PaddleTransactionCompletedHandler struct {
	userRepository                domain.UserRepository
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository
	cachingRepository             domain.CachingRepository
}

func NewPaddleTransactionCompletedHandler(
	userRepository domain.UserRepository,
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository,
	cachingRepository domain.CachingRepository,
) PaddleTransactionCompletedHandler {
	return PaddleTransactionCompletedHandler{
		userRepository:                userRepository,
		subscriptionHistoryRepository: subscriptionHistoryRepository,
		cachingRepository:             cachingRepository,
	}
}

func (h PaddleTransactionCompletedHandler) PaddleTransactionCompleted(ctx *appcontext.AppContext, payload domain.QueuePaddleTransactionCompletedPayload) error {
	tracer := otel.Tracer("tracing")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] Paddle transaction completed")
	ctx.SetContext(spanCtx)
	defer span.End()

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

	ctx.Logger().Text("find subscription in db")
	subscription, err := h.subscriptionHistoryRepository.FindBySourceID(ctx, payload.SubscriptionID)
	if err != nil {
		ctx.Logger().Error("failed to find subscription in db", err, appcontext.Fields{})
		return err
	}
	if subscription == nil {
		ctx.Logger().ErrorText("subscription not found, respond")
		return nil
	}

	ctx.Logger().Text("set subscription activated")
	subscription.SetActivated()

	ctx.Logger().Text("update subscription history in db")
	if err = h.subscriptionHistoryRepository.Update(ctx, *subscription); err != nil {
		ctx.Logger().Error("failed to update subscription history in db", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update user subscription")
	user.SetPlanPro(subscription.NextBilledAt)

	ctx.Logger().Text("update user in db")
	if err = h.userRepository.Update(ctx, *user); err != nil {
		ctx.Logger().Error("failed to update user in db", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("delete user caching data")
	if err = h.cachingRepository.DeleteUserByID(ctx, payload.UserID); err != nil {
		ctx.Logger().Error("failed to delete user caching data", err, appcontext.Fields{})
		return err
	}

	return nil
}
