package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type QueueRepository interface {
	PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload QueuePaddleSubscriptionCreatedPayload) error
	PaddleTransactionCompleted(ctx *appcontext.AppContext, payload QueuePaddleTransactionCompletedPayload) error

	FastspringSubscriptionActivated(ctx *appcontext.AppContext, payload QueueFastspringSubscriptionActivatedPayload) error
}

type QueuePaddleSubscriptionCreatedPayload struct {
	UserID         string
	SubscriptionID string
	NextBilledAt   time.Time
	CustomerID     string
	Items          []string
}

type QueuePaddleTransactionCompletedPayload struct {
	UserID         string
	SubscriptionID string
}

type QueueFastspringSubscriptionActivatedPayload struct {
	UserID         string
	SubscriptionID string
	NextBilledAt   time.Time
	CustomerID     string
	Items          []string
}
