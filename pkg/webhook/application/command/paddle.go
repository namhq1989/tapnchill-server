package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/webhook/dto"
)

type PaddleHandler struct{}

func NewPaddleHandler() PaddleHandler {
	return PaddleHandler{}
}

func (h PaddleHandler) Paddle(ctx *appcontext.AppContext, req dto.PaddleRequest) (*dto.PaddleResponse, error) {
	ctx.Logger().Info("new paddle webhook", appcontext.Fields{"eventType": req.EventType})

	ctx.Logger().Print("payload", req)

	return nil, nil
}
