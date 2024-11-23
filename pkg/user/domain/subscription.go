package domain

import "time"

type Subscription struct {
	Plan   Plan
	Expiry *time.Time
}
