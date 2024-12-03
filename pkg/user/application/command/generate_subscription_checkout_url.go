package command

import (
	"encoding/base64"

	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type GenerateSubscriptionCheckoutURLHandler struct {
	externalAPIRepository domain.ExternalAPIRepository
}

func NewGenerateSubscriptionCheckoutURLHandler(externalAPIRepository domain.ExternalAPIRepository) GenerateSubscriptionCheckoutURLHandler {
	return GenerateSubscriptionCheckoutURLHandler{
		externalAPIRepository: externalAPIRepository,
	}
}

func (h GenerateSubscriptionCheckoutURLHandler) GenerateSubscriptionCheckoutURL(ctx *appcontext.AppContext, performerID string, req dto.GenerateSubscriptionCheckoutURLRequest) (*dto.GenerateSubscriptionCheckoutURLResponse, error) {
	ctx.Logger().Info("new generate subscription checkout url request", appcontext.Fields{"performerID": performerID, "subscriptionID": req.SubscriptionID})

	if req.SubscriptionID != domain.SubscriptionIDMonthly && req.SubscriptionID != domain.SubscriptionIDYearly {
		ctx.Logger().ErrorText("invalid subscription id")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("call external api to generate checkout url")
	checkoutURL, err := h.externalAPIRepository.GenerateLemonsqueezyCheckoutURL(ctx, performerID, req.SubscriptionID)
	if err != nil {
		ctx.Logger().Error("failed to call external api to generate checkout url", err, appcontext.Fields{})
		return nil, err
	}
	if checkoutURL == nil {
		ctx.Logger().ErrorText("checkout url is nil")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("done generate subscription checkout url request")
	return &dto.GenerateSubscriptionCheckoutURLResponse{
		CheckoutURL: base64.StdEncoding.EncodeToString([]byte(*checkoutURL)),
	}, nil
}
