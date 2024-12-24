package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type QueueRepository interface {
	CreateUserDefaultGoal(ctx *appcontext.AppContext, payload QueueCreateUserDefaultGoalPayload) error
}

type QueueCreateUserDefaultGoalPayload struct {
	UserID string
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

type QueueLemonsqueezySubscriptionPaymentSuccessPayload struct {
	UserID    string
	InvoiceID string
}

type QueueDowngradeExpiredSubscriptionsPayload struct{}
