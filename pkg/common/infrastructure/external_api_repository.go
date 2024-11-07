package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/externalapi"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
)

type ExternalAPIRepository struct {
	ea externalapi.Operations
}

func NewExternalAPIRepository(ea externalapi.Operations) ExternalAPIRepository {
	return ExternalAPIRepository{
		ea: ea,
	}
}

func (r ExternalAPIRepository) GetRandomQuote(ctx *appcontext.AppContext) (*domain.Quote, error) {
	apiQuote, err := r.ea.GetRandomQuote(ctx)
	if err != nil {
		return nil, err
	}

	quote, err := domain.NewQuote(apiQuote.OriginalID, apiQuote.Content, apiQuote.Author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
