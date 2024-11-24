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

func (r QueueRepository) SubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueueSubscriptionCreatedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.SubscriptionCreated, payload, 5)
}

func (r QueueRepository) TransactionCompleted(ctx *appcontext.AppContext, payload domain.QueueTransactionCompletedPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.TransactionCompleted, payload, 5)
}
