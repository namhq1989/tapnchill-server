package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
)

func (s Service) GetCityWeather(ctx *appcontext.AppContext, city string) (*domain.Weather, error) {
	ctx.Logger().Info("get city weather", appcontext.Fields{"city": city})

	ctx.Logger().Text("find in caching")
	weather, err := s.cachingRepository.GetCityWeather(ctx, city)
	if weather != nil {
		ctx.Logger().Text("found in caching, respond")
		return weather, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("not found in caching, call api")
	weather, err = s.externalApiRepository.GetCityWeather(ctx, city)
	if err != nil {
		ctx.Logger().Error("failed to call api", err, appcontext.Fields{})
		return nil, err
	}
	if weather == nil {
		ctx.Logger().Text("city weather not found")
		return nil, apperrors.Common.BadRequest
	}

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetCityWeather(ctx, city, *weather); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done get city weather")
	return weather, nil
}
