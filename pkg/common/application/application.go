package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/common/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type (
	Commands interface {
		CreateFeedback(ctx *appcontext.AppContext, performerID string, req dto.CreateFeedbackRequest) (*dto.CreateFeedbackResponse, error)
	}
	Queries interface {
		GetQuote(ctx *appcontext.AppContext, performerID string, _ dto.GetQuoteRequest) (*dto.GetQuoteResponse, error)
		GetWeather(ctx *appcontext.AppContext, performerID string, _ dto.GetWeatherRequest) (*dto.GetWeatherResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.CreateFeedbackHandler
	}
	queryHandlers struct {
		query.GetQuoteHandler
		query.GetWeatherHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	feedbackRepository domain.FeedbackRepository,
	quoteRepository domain.QuoteRepository,
	service domain.Service,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateFeedbackHandler: command.NewCreateFeedbackHandler(feedbackRepository),
		},
		queryHandlers: queryHandlers{
			GetQuoteHandler:   query.NewGetQuoteHandler(quoteRepository),
			GetWeatherHandler: query.NewGetWeatherHandler(service),
		},
	}
}
