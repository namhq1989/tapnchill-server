package dto

import "github.com/namhq1989/tapnchill-server/pkg/common/domain"

type Weather struct {
	Temp       float64 `json:"temp"`
	FeelsLike  float64 `json:"feelsLike"`
	Humidity   float64 `json:"humidity"`
	WindSpeed  float64 `json:"windSpeed"`
	PrecipProb float64 `json:"precipProb"`
	Icon       string  `json:"icon"`
}

func (Weather) FromDomain(weather domain.Weather) Weather {
	return Weather{
		Temp:       weather.Temp,
		FeelsLike:  weather.FeelsLike,
		Humidity:   weather.Humidity,
		WindSpeed:  weather.WindSpeed,
		PrecipProb: weather.PrecipProb,
		Icon:       weather.Icon,
	}
}
