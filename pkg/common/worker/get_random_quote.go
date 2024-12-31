package worker

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"go.opentelemetry.io/otel"
)

type GetRandomQuoteHandler struct {
	quoteRepository       domain.QuoteRepository
	externalApiRepository domain.ExternalApiRepository
}

func NewGetRandomQuoteHandler(quoteRepository domain.QuoteRepository, externalApiRepository domain.ExternalApiRepository) GetRandomQuoteHandler {
	return GetRandomQuoteHandler{
		quoteRepository:       quoteRepository,
		externalApiRepository: externalApiRepository,
	}
}

func (h GetRandomQuoteHandler) GetRandomQuote(ctx *appcontext.AppContext, _ domain.QueueGetRandomQuotePayload) error {
	tracer := otel.Tracer("tracing")
	spanCtx, span := tracer.Start(ctx.Context(), "[worker] get random quote")
	ctx.SetContext(spanCtx)
	defer span.End()

	ctx.Logger().Text("call api to get random quote")
	quote, err := h.externalApiRepository.GetRandomQuote(ctx)
	if err != nil {
		ctx.Logger().Error("failed to call api to get random quote", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist quote in db")
	if err = h.quoteRepository.Create(ctx, *quote); err != nil {
		ctx.Logger().Error("failed to persist quote in db", err, appcontext.Fields{})
		return err
	}

	return nil
}
