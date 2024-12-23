package query

import (
	"os"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

var resourceLimitations = dto.ResourcesLimitation{
	Goal: dto.PlanLimitation{
		Free: domain.FreePlanMaxGoals,
		Pro:  domain.ProPlanMaxGoals,
	},
	Task: dto.PlanLimitation{
		Free: domain.FreePlanMaxTaskPerGoal,
		Pro:  domain.ProPlanMaxTaskPerGoal,
	},
	Habit: dto.PlanLimitation{
		Free: domain.FreePlanMaxHabits,
		Pro:  domain.ProPlanMaxHabits,
	},
	Note: dto.PlanLimitation{
		Free: domain.FreePlanMaxNotes,
		Pro:  domain.ProPlanMaxNotes,
	},
	QRCode: dto.PlanLimitation{
		Free: domain.FreePlanMaxQRCodes,
		Pro:  domain.ProPlanMaxQRCodes,
	},
}

type GetSubscriptionPlansHandler struct{}

func NewGetSubscriptionPlansHandler() GetSubscriptionPlansHandler {
	return GetSubscriptionPlansHandler{}
}

func (h GetSubscriptionPlansHandler) GetSubscriptionPlans(ctx *appcontext.AppContext, performerID string, _ dto.GetSubscriptionPlansRequest) (*dto.GetSubscriptionPlansResponse, error) {
	ctx.Logger().Info("new get subscription plans request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("detect whether subscription is enabled or not")
	isSubscriptionEnabled := false

	// ctx.Logger().Print("IS_SUBSCRIPTION_ENABLED", os.Getenv("IS_SUBSCRIPTION_ENABLED"))
	if os.Getenv("IS_SUBSCRIPTION_ENABLED") == "true" {
		isSubscriptionEnabled = true
	}

	if !isSubscriptionEnabled {
		ctx.Logger().Text("subscription is disabled, respond")
		return &dto.GetSubscriptionPlansResponse{
			IsEnabled:           false,
			Plans:               []dto.SubscriptionPlan{},
			ResourcesLimitation: resourceLimitations,
		}, nil
	}

	ctx.Logger().Text("prepare data")

	monthly := dto.SubscriptionPlan{
		ID:                  domain.SubscriptionIDMonthly,
		Amount:              3,
		AfterDiscountAmount: 3,
	}

	yearly := dto.SubscriptionPlan{
		ID:                  domain.SubscriptionIDYearly,
		Amount:              36,
		AfterDiscountAmount: 30,
	}

	ctx.Logger().Text("done get subscription plans request")
	return &dto.GetSubscriptionPlansResponse{
		IsEnabled: true,
		Plans: []dto.SubscriptionPlan{
			monthly,
			yearly,
		},
		ResourcesLimitation: resourceLimitations,
	}, nil
}
