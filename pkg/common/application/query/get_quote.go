package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type GetQuoteHandler struct {
	quoteRepository domain.QuoteRepository
}

func NewGetQuoteHandler(quoteRepository domain.QuoteRepository) GetQuoteHandler {
	return GetQuoteHandler{
		quoteRepository: quoteRepository,
	}
}

func (h GetQuoteHandler) GetQuote(ctx *appcontext.AppContext, performerID string, _ dto.GetQuoteRequest) (*dto.GetQuoteResponse, error) {
	ctx.Logger().Info("new get quote request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("find latest quote in db")
	quote, err := h.quoteRepository.FindLatest(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find latest quote in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done get quote request")
	return &dto.GetQuoteResponse{
		Quote: dto.Quote{}.FromDomain(*quote),
	}, nil
}
