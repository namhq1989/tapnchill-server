package dto

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type UserSubscription struct {
	Plan       string                    `json:"plan"`
	Expiry     *httprespond.TimeResponse `json:"expiry"`
	CustomerID string                    `json:"customerId"`
}

func (UserSubscription) FromDomain(subscription domain.UserSubscription) UserSubscription {
	s := UserSubscription{
		Plan:       string(subscription.Plan),
		CustomerID: subscription.CustomerID,
	}

	if subscription.Expiry != nil {
		s.Expiry = httprespond.NewTimeResponse(*subscription.Expiry)
	} else {
		s.Expiry = httprespond.NewTimeResponse(time.Time{})
	}

	return s
}
