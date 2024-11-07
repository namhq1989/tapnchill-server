package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type GetQuoteHandler struct {
	quoteRepository   domain.QuoteRepository
	cachingRepository domain.CachingRepository
}

func NewGetQuoteHandler(quoteRepository domain.QuoteRepository, cachingRepository domain.CachingRepository) GetQuoteHandler {
	return GetQuoteHandler{
		quoteRepository:   quoteRepository,
		cachingRepository: cachingRepository,
	}
}

func (h GetQuoteHandler) GetQuote(ctx *appcontext.AppContext, performerID string, _ dto.GetQuoteRequest) (*dto.GetQuoteResponse, error) {
	ctx.Logger().Info("new get quote request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("find in caching")
	quote, err := h.cachingRepository.GetLatestQuote(ctx)
	if quote != nil {
		ctx.Logger().Text("found in caching, respond")
		return &dto.GetQuoteResponse{
			Quote: dto.Quote{}.FromDomain(*quote),
		}, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find latest quote in db")
	quote, err = h.quoteRepository.FindLatest(ctx)
	if err != nil {
		ctx.Logger().Error("failed to find latest quote in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist in caching")
	if err = h.cachingRepository.SetLatestQuote(ctx, *quote); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("done get quote request")
	return &dto.GetQuoteResponse{
		Quote: dto.Quote{}.FromDomain(*quote),
	}, nil
}
