package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type (
	Handlers interface {
		PaddleSubscriptionCreated(ctx *appcontext.AppContext, payload domain.QueuePaddleSubscriptionCreatedPayload) error
		PaddleTransactionCompleted(ctx *appcontext.AppContext, payload domain.QueuePaddleTransactionCompletedPayload) error

		LemonsqueezySubscriptionPaymentSuccess(ctx *appcontext.AppContext, payload domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload) error
	}
	Cronjob interface {
		DowngradeExpiredSubscriptions(ctx *appcontext.AppContext, _ domain.QueueDowngradeExpiredSubscriptionsPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		PaddleSubscriptionCreatedHandler
		PaddleTransactionCompletedHandler

		LemonsqueezySubscriptionPaymentSuccessHandler
	}
	workerCronjob struct {
		DowngradeExpiredSubscriptionsHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
		workerCronjob
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	userRepository domain.UserRepository,
	subscriptionHistoryRepository domain.SubscriptionHistoryRepository,
	cachingRepository domain.CachingRepository,
	externalAPIRepository domain.ExternalAPIRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			PaddleSubscriptionCreatedHandler:  NewPaddleSubscriptionCreatedHandler(userRepository, subscriptionHistoryRepository),
			PaddleTransactionCompletedHandler: NewPaddleTransactionCompletedHandler(userRepository, subscriptionHistoryRepository, cachingRepository),

			LemonsqueezySubscriptionPaymentSuccessHandler: NewLemonsqueezySubscriptionPaymentSuccessHandler(userRepository, subscriptionHistoryRepository, cachingRepository, externalAPIRepository),
		},
		workerCronjob: workerCronjob{
			DowngradeExpiredSubscriptionsHandler: NewDowngradeExpiredSubscriptionsHandler(userRepository),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.PaddleSubscriptionCreated), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueuePaddleSubscriptionCreatedPayload](bgCtx, t, queue.ParsePayload[domain.QueuePaddleSubscriptionCreatedPayload], w.PaddleSubscriptionCreated)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.PaddleTransactionCompleted), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueuePaddleTransactionCompletedPayload](bgCtx, t, queue.ParsePayload[domain.QueuePaddleTransactionCompletedPayload], w.PaddleTransactionCompleted)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.LemonsqueezySubscriptionPaymentSuccess), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload](bgCtx, t, queue.ParsePayload[domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload], w.LemonsqueezySubscriptionPaymentSuccess)
	})

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.DowngradeExpiredSubscriptions), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueDowngradeExpiredSubscriptionsPayload](bgCtx, t, queue.ParsePayload[domain.QueueDowngradeExpiredSubscriptionsPayload], w.DowngradeExpiredSubscriptions)
	})
}

type cronjobData struct {
	Task       string      `json:"task"`
	CronSpec   string      `json:"cronSpec"`
	Payload    interface{} `json:"payload"`
	RetryTimes int         `json:"retryTimes"`
}

func (w Worker) addCronjob() {
	var (
		ctx  = appcontext.NewWorker(context.Background())
		jobs = []cronjobData{
			{
				Task:       w.queue.GenerateTypename(queue.TypeNames.DowngradeExpiredSubscriptions),
				CronSpec:   "1 0 * * *", // 00:01 AM everyday
				Payload:    domain.QueueDowngradeExpiredSubscriptionsPayload{},
				RetryTimes: 3,
			},
		}
	)

	for _, job := range jobs {
		entryID, err := w.queue.ScheduleTask(job.Task, job.Payload, job.CronSpec, job.RetryTimes)
		if err != nil {
			ctx.Logger().Error("error when initializing cronjob", err, appcontext.Fields{"job": job})
			panic(err)
		}

		ctx.Logger().Info(fmt.Sprintf("[cronjob] cronjob '%s' initialize successfully with cronSpec '%s' and retryTimes '%d'", job.Task, job.CronSpec, job.RetryTimes), appcontext.Fields{
			"entryId": entryID,
		})
	}
}
