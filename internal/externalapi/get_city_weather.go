package externalapi

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
)

type GetCityWeatherResult struct {
	Temp       float64
	FeelsLike  float64
	Humidity   float64
	WindSpeed  float64
	PrecipProb float64
	Icon       string
}

type getCityWeatherApiResult struct {
	CurrentConditions getCityWeatherApiResultCurrent `json:"currentConditions"`
}

type getCityWeatherApiResultCurrent struct {
	Temp       float64 `json:"temp"`
	FeelsLike  float64 `json:"feelslike"`
	Humidity   float64 `json:"humidity"`
	WindSpeed  float64 `json:"windspeed"`
	PrecipProb float64 `json:"precipprob"`
	Icon       string  `json:"icon"`
}

func (ea ExternalApi) GetCityWeather(ctx *appcontext.AppContext, city string) (*GetCityWeatherResult, error) {
	var (
		apiResult = &getCityWeatherApiResult{}
	)

	_, err := ea.weatherClient.R().
		SetQueryParams(map[string]string{
			"unitGroup":   "metric",
			"key":         ea.visualCrossingToken,
			"contentType": "json",
		}).
		SetResult(&apiResult).
		Get(fmt.Sprintf("/VisualCrossingWebServices/rest/services/timeline/%s", city))

	if err != nil {
		ctx.Logger().Error("[externalapi] error when get city by ip", err, appcontext.Fields{})
		return nil, err
	}

	if apiResult == nil || apiResult.CurrentConditions.Icon == "" {
		return nil, nil
	}

	return &GetCityWeatherResult{
		Temp:       apiResult.CurrentConditions.Temp,
		FeelsLike:  apiResult.CurrentConditions.FeelsLike,
		Humidity:   apiResult.CurrentConditions.Humidity,
		WindSpeed:  apiResult.CurrentConditions.WindSpeed,
		PrecipProb: apiResult.CurrentConditions.PrecipProb,
		Icon:       apiResult.CurrentConditions.Icon,
	}, nil
}
