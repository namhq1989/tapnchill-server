package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type (
	Commands interface {
		CreateFeedback(ctx *appcontext.AppContext, performerID string, req dto.CreateFeedbackRequest) (*dto.CreateFeedbackResponse, error)
	}
	Instance interface {
		Commands
	}

	commandHandlers struct {
		command.CreateFeedbackHandler
	}
	Application struct {
		commandHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	feedbackRepository domain.FeedbackRepository,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateFeedbackHandler: command.NewCreateFeedbackHandler(feedbackRepository),
		},
	}
}
