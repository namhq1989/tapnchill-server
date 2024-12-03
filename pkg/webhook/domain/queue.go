package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
)

type QueueRepository interface {
	PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload QueuePaddleSubscriptionCreatedPayload) error
	PaddleTransactionCompleted(ctx *appcontext.AppContext, payload QueuePaddleTransactionCompletedPayload) error

	LemonsqueezySubscriptionPaymentSuccess(ctx *appcontext.AppContext, payload QueueLemonsqueezySubscriptionPaymentSuccessPayload) error
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
