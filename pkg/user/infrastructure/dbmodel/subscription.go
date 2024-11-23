package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type Subscription struct {
	Plan   string     `bson:"plan"`
	Expiry *time.Time `bson:"expiry"`
}

func (s Subscription) ToDomain() domain.Subscription {
	return domain.Subscription{
		Plan:   domain.ToPlan(s.Plan),
		Expiry: s.Expiry,
	}
}

func (Subscription) FromDomain(subscription domain.Subscription) Subscription {
	return Subscription{
		Plan:   subscription.Plan.String(),
		Expiry: subscription.Expiry,
	}
}
