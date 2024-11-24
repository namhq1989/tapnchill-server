package domain

import "time"

type UserSubscription struct {
	Plan       Plan
	Expiry     *time.Time
	CustomerID string
}
