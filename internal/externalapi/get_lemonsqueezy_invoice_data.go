package externalapi

import (
	"fmt"
	"strconv"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type GetLemonsqueezySubscriptionDataResult struct {
	SubscriptionID string
	CustomerID     string
	VariantID      string
	RenewsAt       *time.Time
}

type getLemonsqueezyInvoiceDataResponse struct {
	Data getLemonsqueezyInvoiceDataResponseData `json:"data"`
}

type getLemonsqueezyInvoiceDataResponseData struct {
	ID         string                                           `json:"id"`
	Attributes getLemonsqueezyInvoiceDataResponseDataAttributes `json:"attributes"`
}

type getLemonsqueezyInvoiceDataResponseDataAttributes struct {
	CustomerID int    `json:"customer_id"`
	VariantID  int    `json:"variant_id"`
	RenewsAt   string `json:"renews_at"`
}

func (ea ExternalApi) GetLemonsqueezySubscriptionInvoiceData(ctx *appcontext.AppContext, invoiceID string) (*GetLemonsqueezySubscriptionDataResult, error) {
	var apiResult = &getLemonsqueezyInvoiceDataResponse{}

	_, err := ea.lemonsqueezyClient.R().
		SetResult(&apiResult).
		Get(fmt.Sprintf("/v1/subscription-invoices/%s/subscription", invoiceID))

	if err != nil {
		ctx.Logger().Error("[externalapi] error when get Lemonsqueezy subscription data", err, appcontext.Fields{})
		return nil, err
	}

	t, _ := manipulation.GetEndOfDayWithClientDate(apiResult.Data.Attributes.RenewsAt)
	return &GetLemonsqueezySubscriptionDataResult{
		SubscriptionID: apiResult.Data.ID,
		CustomerID:     strconv.Itoa(apiResult.Data.Attributes.CustomerID),
		VariantID:      strconv.Itoa(apiResult.Data.Attributes.VariantID),
		RenewsAt:       t,
	}, nil
}
