package dto

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type Subscription struct {
	Plan   string                    `json:"plan"`
	Expiry *httprespond.TimeResponse `json:"expiry"`
}

func (Subscription) FromDomain(subscription domain.Subscription) Subscription {
	s := Subscription{
		Plan: string(subscription.Plan),
	}

	if subscription.Expiry != nil {
		s.Expiry = httprespond.NewTimeResponse(*subscription.Expiry)
	} else {
		s.Expiry = httprespond.NewTimeResponse(time.Time{})
	}

	return s
}
