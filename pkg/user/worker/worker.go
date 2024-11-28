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
		PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueuePaddleSubscriptionCreatedPayload) error
		PaddleTransactionCompleted(ctx *appcontext.AppContext, payload domain.QueuePaddleTransactionCompletedPayload) error

		FastspringSubscriptionActivated(ctx *appcontext.AppContext, payload domain.QueueFastspringSubscriptionActivatedPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		PaddleSubscriptionCreatedHandler
		PaddleTransactionCompletedHandler

		FastspringSubscriptionActivatedHandler
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
			PaddleSubscriptionCreatedHandler:  NewPaddleSubscriptionCreatedHandler(userRepository, subscriptionHistoryRepository),
			PaddleTransactionCompletedHandler: NewPaddleTransactionCompletedHandler(userRepository, subscriptionHistoryRepository, cachingRepository),

			FastspringSubscriptionActivatedHandler: NewFastspringSubscriptionActivatedHandler(userRepository, subscriptionHistoryRepository, cachingRepository),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.PaddleSubscriptionCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueuePaddleSubscriptionCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueuePaddleSubscriptionCreatedPayload], w.PaddleSubscriptionCreated)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.PaddleTransactionCompleted), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueuePaddleTransactionCompletedPayload](bgCtx, t, queue.ParsePayload[domain.QueuePaddleTransactionCompletedPayload], w.PaddleTransactionCompleted)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.FastspringSubscriptionActivated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueFastspringSubscriptionActivatedPayload](bgCtx, t, queue.ParsePayload[domain.QueueFastspringSubscriptionActivatedPayload], w.FastspringSubscriptionActivated)
	})
}
