package worker

import (
	"os"

	"go.opentelemetry.io/otel"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type LemonsqueezySubscriptionPaymentSuccessHandler struct {
	userRepository                domain.UserRepository
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository
	cachingRepository             domain.CachingRepository
	externalAPIRepository         domain.ExternalAPIRepository
}

func NewLemonsqueezySubscriptionPaymentSuccessHandler(
	userRepository domain.UserRepository,
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository,
	cachingRepository domain.CachingRepository,
	externalAPIRepository domain.ExternalAPIRepository,
) LemonsqueezySubscriptionPaymentSuccessHandler {
	return LemonsqueezySubscriptionPaymentSuccessHandler{
		userRepository:                userRepository,
		subscriptionHistoryRepository: subscriptionHistoryRepository,
		cachingRepository:             cachingRepository,
		externalAPIRepository:         externalAPIRepository,
	}
}

func (h LemonsqueezySubscriptionPaymentSuccessHandler) LemonsqueezySubscriptionPaymentSuccess(ctx *appcontext.AppContext, payload domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload) error {
	tracer := otel.Tracer("tracing")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] LemonSqueezy subscription payment success")
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

	ctx.Logger().Text("get Lemonsqueezy invoice data")
	invoiceData, err := h.externalAPIRepository.GetLemonsqueezyInvoiceData(ctx, payload.InvoiceID)
	if err != nil {
		ctx.Logger().Error("failed to get invoice data", err, appcontext.Fields{})
		return err
	}
	if invoiceData == nil {
		ctx.Logger().ErrorText("invoice data not found, respond")
		return nil
	}

	ctx.Logger().Info("got invoice data", appcontext.Fields{
		"subscriptionID": invoiceData.SubscriptionID, "variantID": invoiceData.VariantID, "renewsAt": invoiceData.RenewsAt,
	})

	nextBilledAt := invoiceData.RenewsAt
	if nextBilledAt == nil {
		ctx.Logger().Text("renews at not found, calculate next billed at")
		now := manipulation.NowUTC()
		if invoiceData.VariantID == os.Getenv("LEMONSQUEEZY_SUBSCRIPTION_YEARLY_VARIANT_ID") {
			now = now.AddDate(1, 0, 0)
		} else {
			now = now.AddDate(0, 1, 0)
		}
		now = manipulation.EndOfDay(now)
		nextBilledAt = &now
	}

	ctx.Logger().Text("create new subscription history model")
	history, err := domain.NewSubscriptionHistory(payload.UserID, invoiceData.SubscriptionID, domain.PaymentSourceLemonsqueezy, invoiceData.CustomerID, []string{invoiceData.VariantID}, *nextBilledAt)
	if err != nil {
		ctx.Logger().Error("failed to create new subscription history model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("set subscription activated")
	history.SetActivated()

	ctx.Logger().Text("persist subscription history in db")
	if err = h.subscriptionHistoryRepository.Create(ctx, *history); err != nil {
		ctx.Logger().Error("failed to persist subscription history in db", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("update user subscription")
	user.SetPlanPro(history.NextBilledAt)
	user.SetSubscriptionCustomerID(history.SourceCustomerID)

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
