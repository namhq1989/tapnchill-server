package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type QueueRepository interface {
	SubscriptionCreated(ctx *appcontext.AppContext, payload QueueSubscriptionCreatedPayload) error
	TransactionCompleted(ctx *appcontext.AppContext, payload QueueTransactionCompletedPayload) error
}

type QueueSubscriptionCreatedPayload struct {
	UserID         string
	SubscriptionID string
	NextBilledAt   time.Time
	CustomerID     string
	Items          []string
}

type QueueTransactionCompletedPayload struct {
	UserID         string
	SubscriptionID string
}
