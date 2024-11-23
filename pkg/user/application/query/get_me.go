package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
)

type GetMeHandler struct {
	serviceRepository domain.Service
}

func NewGetMeHandler(serviceRepository domain.Service) GetMeHandler {
	return GetMeHandler{
		serviceRepository: serviceRepository,
	}
}

func (h GetMeHandler) GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error) {
	ctx.Logger().Print("new get me request", appcontext.Fields{"performerID": performerID})

	user, err := h.serviceRepository.GetUserByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find user in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("done get me request")
	return &dto.GetMeResponse{
		Ip:           ctx.GetIP(),
		Subscription: dto.Subscription{}.FromDomain(user.Subscription),
	}, nil
}
