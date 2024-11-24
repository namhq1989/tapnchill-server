package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type (
	Handlers interface {
		SubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueueSubscriptionCreatedPayload) error
		TransactionCompleted(ctx *appcontext.AppContext, payload domain.QueueTransactionCompletedPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		SubscriptionCreatedHandler
		TransactionCompletedHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	userRepository domain.UserRepository,
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository,
	cachingRepository domain.CachingRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			SubscriptionCreatedHandler:  NewSubscriptionCreatedHandler(userRepository, subscriptionHistoryRepository),
			TransactionCompletedHandler: NewTransactionCompletedHandler(userRepository, subscriptionHistoryRepository, cachingRepository),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.SubscriptionCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueSubscriptionCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueueSubscriptionCreatedPayload], w.SubscriptionCreated)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.TransactionCompleted), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueTransactionCompletedPayload](bgCtx, t, queue.ParsePayload[domain.QueueTransactionCompletedPayload], w.TransactionCompleted)
	})
}
