package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
)

func (s Service) GetTaskByID(ctx *appcontext.AppContext, taskID, userID string) (*domain.Task, error) {
	ctx.Logger().Text("find task in db")
	task, err := s.taskRepository.FindByID(ctx, taskID)
	if err != nil {
		ctx.Logger().Error("failed to find task in db", err, appcontext.Fields{})
		return nil, err
	}
	if task == nil {
		ctx.Logger().ErrorText("task not found, respond")
		return nil, apperrors.Common.NotFound
	}

	if task.UserID != userID {
		ctx.Logger().ErrorText("task author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	return task, nil
}
