package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/queue"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type QueueRepository struct {
	queue queue.Operations
}

func NewQueueRepository(queue queue.Operations) QueueRepository {
	return QueueRepository{
		queue: queue,
	}
}

func (r QueueRepository) CreateUserDefaultGoal(ctx *appcontext.AppContext, payload domain.QueueCreateUserDefaultGoalPayload) error {
	return queue.EnqueueTask(ctx, r.queue, queue.TypeNames.CreateUserDefaultGoal, payload, -1)
}
