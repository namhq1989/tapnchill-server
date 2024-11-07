package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type GetWeatherHandler struct {
	service domain.Service
}

func NewGetWeatherHandler(service domain.Service) GetWeatherHandler {
	return GetWeatherHandler{
		service: service,
	}
}

func (h GetWeatherHandler) GetWeather(ctx *appcontext.AppContext, performerID string, _ dto.GetWeatherRequest) (*dto.GetWeatherResponse, error) {
	ctx.Logger().Info("new get weather request", appcontext.Fields{"performerID": performerID, "ip": ctx.GetIP()})

	ctx.Logger().Text("get ip city")
	city, err := h.service.GetIpCity(ctx, ctx.GetIP())
	if err != nil {
		ctx.Logger().Error("failed to get ip city", err, appcontext.Fields{})
		return nil, err
	}
	if city == nil {
		ctx.Logger().ErrorText("ip city not found")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("get city weather")
	weather, err := h.service.GetCityWeather(ctx, *city)
	if err != nil {
		ctx.Logger().Error("failed to get city weather", err, appcontext.Fields{})
		return nil, err
	}
	if weather == nil {
		ctx.Logger().ErrorText("city weather not found")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("done get weather request")
	return &dto.GetWeatherResponse{
		City:    *city,
		Weather: dto.Weather{}.FromDomain(*weather),
	}, nil
}
