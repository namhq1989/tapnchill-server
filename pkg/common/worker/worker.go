package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
)

type (
	Cronjob interface {
		GetRandomQuote(ctx *appcontext.AppContext, _ domain.QueueGetRandomQuotePayload) error
	}
	Instance interface {
		Cronjob
	}

	workerCronjob struct {
		GetRandomQuoteHandler
	}
	Worker struct {
		queue queue.Operations
		workerCronjob
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	quoteRepository domain.QuoteRepository,
	externalApiRepository domain.ExternalApiRepository,
) Worker {
	return Worker{
		queue: queue,
		workerCronjob: workerCronjob{
			GetRandomQuoteHandler: NewGetRandomQuoteHandler(quoteRepository, externalApiRepository),
		},
	}
}

func (w Worker) Start() {
	w.addCronjob()

	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.GetRandomQuote), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueGetRandomQuotePayload](bgCtx, t, queue.ParsePayload[domain.QueueGetRandomQuotePayload], w.GetRandomQuote)
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
				Task:       w.queue.GenerateTypename(queue.TypeNames.GetRandomQuote),
				CronSpec:   "0 */3 * * *", // every 3h
				Payload:    domain.QueueGetRandomQuotePayload{},
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
