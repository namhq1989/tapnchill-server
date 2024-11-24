package query

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type GetSubscriptionPlansHandler struct{}

func NewGetSubscriptionPlansHandler() GetSubscriptionPlansHandler {
	return GetSubscriptionPlansHandler{}
}

func (h GetSubscriptionPlansHandler) GetSubscriptionPlans(ctx *appcontext.AppContext, performerID string, _ dto.GetSubscriptionPlansRequest) (*dto.GetSubscriptionPlansResponse, error) {
	ctx.Logger().Info("new get subscription plans request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("prepare data")
	monthly := dto.SubscriptionPlan{
		PeriodText:          "month",
		PriceID:             os.Getenv("SUBSCRIPTION_MONTHLY_PRICE_ID"),
		Amount:              3,
		DiscountID:          "",
		AfterDiscountAmount: 3,
		Token:               base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", performerID, os.Getenv("SUBSCRIPTION_MONTHLY_PRICE_ID")))),
	}

	yearly := dto.SubscriptionPlan{
		PeriodText:          "year",
		PriceID:             os.Getenv("SUBSCRIPTION_YEARLY_PRICE_ID"),
		Amount:              36,
		DiscountID:          os.Getenv("SUBSCRIPTION_YEARLY_DISCOUNT_ID"),
		AfterDiscountAmount: 30,
		Token:               base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s:%s", performerID, os.Getenv("SUBSCRIPTION_YEARLY_PRICE_ID"), os.Getenv("SUBSCRIPTION_YEARLY_DISCOUNT_ID")))),
	}

	ctx.Logger().Text("done get subscription plans request")
	return &dto.GetSubscriptionPlansResponse{
		Plans: []dto.SubscriptionPlan{monthly, yearly},
	}, nil
}
