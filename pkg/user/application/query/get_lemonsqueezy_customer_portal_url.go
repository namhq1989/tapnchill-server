package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type GetPaymentCustomerPortalURLHandler struct {
	service domain.Service
}

func NewGetPaymentCustomerPortalURLHandler(service domain.Service) GetPaymentCustomerPortalURLHandler {
	return GetPaymentCustomerPortalURLHandler{
		service: service,
	}
}

func (h GetPaymentCustomerPortalURLHandler) GetPaymentCustomerPortalURL(ctx *appcontext.AppContext, performerID string, _ dto.GetPaymentCustomerPortalURLRequest) (*dto.GetPaymentCustomerPortalURLResponse, error) {
	ctx.Logger().Info("new get payment [lemonsqueezy] customer portal url request", appcontext.Fields{"performerID": performerID})

	url, err := h.service.GetLemonsqueezyCustomerPortalURL(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get lemonsqueezy customer portal url", err, appcontext.Fields{})
		return nil, err
	}
	if url == nil {
		ctx.Logger().ErrorText("lemonsqueezy customer portal url is nil")
		return nil, apperrors.Common.SomethingWentWrong
	}

	ctx.Logger().Text("done get lemonsqueezy customer portal url request")
	return &dto.GetPaymentCustomerPortalURLResponse{
		URL: *url,
	}, nil
}
