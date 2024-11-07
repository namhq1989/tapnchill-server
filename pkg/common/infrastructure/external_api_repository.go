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
	apiResult, err := r.ea.GetRandomQuote(ctx)
	if err != nil {
		return nil, err
	}

	quote, err := domain.NewQuote(apiResult.OriginalID, apiResult.Content, apiResult.Author)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (r ExternalAPIRepository) GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error) {
	apiResult, err := r.ea.GetIpCity(ctx, ip)
	if err != nil {
		return nil, err
	}
	if apiResult == nil {
		return nil, nil
	}

	return &apiResult.City, nil
}

func (r ExternalAPIRepository) GetCityWeather(ctx *appcontext.AppContext, city string) (*domain.Weather, error) {
	apiResult, err := r.ea.GetCityWeather(ctx, city)
	if err != nil {
		return nil, err
	}
	if apiResult == nil {
		return nil, nil
	}

	return &domain.Weather{
		Temp:       apiResult.Temp,
		FeelsLike:  apiResult.FeelsLike,
		Humidity:   apiResult.Humidity,
		WindSpeed:  apiResult.WindSpeed,
		PrecipProb: apiResult.PrecipProb,
		Icon:       apiResult.Icon,
	}, nil
}
