package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type UserSubscription struct {
	Plan       string     `bson:"plan"`
	Expiry     *time.Time `bson:"expiry"`
	CustomerID string     `bson:"customerId"`
}

func (s UserSubscription) ToDomain() domain.UserSubscription {
	return domain.UserSubscription{
		Plan:       domain.ToPlan(s.Plan),
		Expiry:     s.Expiry,
		CustomerID: s.CustomerID,
	}
}

func (UserSubscription) FromDomain(subscription domain.UserSubscription) UserSubscription {
	return UserSubscription{
		Plan:       subscription.Plan.String(),
		Expiry:     subscription.Expiry,
		CustomerID: subscription.CustomerID,
	}
}
