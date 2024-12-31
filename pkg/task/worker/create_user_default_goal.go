package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"go.opentelemetry.io/otel"
)

type CreateUserDefaultGoalHandler struct {
	goalRepository domain.GoalRepository
}

func NewCreateUserDefaultGoalHandler(goalRepository domain.GoalRepository) CreateUserDefaultGoalHandler {
	return CreateUserDefaultGoalHandler{
		goalRepository: goalRepository,
	}
}

func (h CreateUserDefaultGoalHandler) CreateUserDefaultGoal(ctx *appcontext.AppContext, payload domain.QueueCreateUserDefaultGoalPayload) error {
	tracer := otel.Tracer("tracing")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] create user default goal")
	ctx.SetContext(spanCtx)
	defer span.End()

	ctx.Logger().Text("create user default goal")
	goal, err := domain.NewGoal(payload.UserID, domain.DefaultGoalName, domain.DefaultGoalDescription)
	if err != nil {
		ctx.Logger().Error("failed to create user default goal", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist goal in db")
	if err = h.goalRepository.Create(ctx, *goal); err != nil {
		ctx.Logger().Error("failed to persist user default goal in db", err, appcontext.Fields{})
		return err
	}

	return nil
}
