package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
)

type (
	Handlers interface {
		CreateUserDefaultGoal(ctx *appcontext.AppContext, payload domain.QueueCreateUserDefaultGoalPayload) error
	}
	Instance interface {
		Handlers
	}

	workerHandlers struct {
		CreateUserDefaultGoalHandler
	}
	Worker struct {
		queue queue.Operations
		workerHandlers
	}
)

var _ Instance = (*Worker)(nil)

func New(
	queue queue.Operations,
	goalRepository domain.GoalRepository,
) Worker {
	return Worker{
		queue: queue,
		workerHandlers: workerHandlers{
			CreateUserDefaultGoalHandler: NewCreateUserDefaultGoalHandler(goalRepository),
		},
	}
}

func (w Worker) Start() {
	server := w.queue.GetServer()

	server.HandleFunc(w.queue.GenerateTypename(queue.TypeNames.CreateUserDefaultGoal), func(bgCtx context.Context, t *asynq.Task) error {
		return queue.ProcessTask[domain.QueueCreateUserDefaultGoalPayload](bgCtx, t, queue.ParsePayload[domain.QueueCreateUserDefaultGoalPayload], w.CreateUserDefaultGoal)
	})
}
