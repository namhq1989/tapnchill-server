package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/domain"
)

type QueueRepository struct {
	queue queue.Operations
}

func NewQueueRepository(queue queue.Operations) QueueRepository {
	return QueueRepository{
		queue: queue,
	}
}

func (r QueueRepository) PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueuePaddleSubscriptionCreatedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.PaddleSubscriptionCreated, payload, 5)
}

func (r QueueRepository) PaddleTransactionCompleted(ctx *appcontext.AppContext, payload domain.QueuePaddleTransactionCompletedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.PaddleTransactionCompleted, payload, 5)
}

func (r QueueRepository) LemonsqueezySubscriptionPaymentSuccess(ctx *appcontext.AppContext, payload domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.LemonsqueezySubscriptionPaymentSuccess, payload, 5)
}
