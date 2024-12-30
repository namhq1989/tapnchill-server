package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/externalapi"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type ExternalAPIRepository struct {
	ea externalapi.Operations
}

func NewExternalAPIRepository(ea externalapi.Operations) ExternalAPIRepository {
	return ExternalAPIRepository{
		ea: ea,
	}
}

func (r ExternalAPIRepository) GenerateLemonsqueezyCheckoutURL(ctx *appcontext.AppContext, userID, subscriptionID string) (*string, error) {
	return r.ea.GenerateLemonsqueezySubscriptionCheckoutURL(ctx, userID, subscriptionID)
}

func (r ExternalAPIRepository) GetLemonsqueezyInvoiceData(ctx *appcontext.AppContext, invoiceID string) (*domain.GetLemonsqueezyInvoiceDataResult, error) {
	resp, err := r.ea.GetLemonsqueezySubscriptionInvoiceData(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	return &domain.GetLemonsqueezyInvoiceDataResult{
		SubscriptionID: resp.SubscriptionID,
		CustomerID:     resp.CustomerID,
		VariantID:      resp.VariantID,
		RenewsAt:       resp.RenewsAt,
	}, nil
}

func (r ExternalAPIRepository) GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, customerID string) (*string, error) {
	resp, err := r.ea.GetLemonsqueezyCustomerPortalURL(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return &resp.URL, nil
}
