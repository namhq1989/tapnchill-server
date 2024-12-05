package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type ExternalAPIRepository interface {
	GenerateLemonsqueezyCheckoutURL(ctx *appcontext.AppContext, userID, subscriptionID string) (*string, error)
	GetLemonsqueezyInvoiceData(ctx *appcontext.AppContext, invoiceID string) (*GetLemonsqueezyInvoiceDataResult, error)
}

type GetLemonsqueezyInvoiceDataResult struct {
	SubscriptionID string
	CustomerID     string
	VariantID      string
	RenewsAt       *time.Time
}
