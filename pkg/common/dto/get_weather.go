package dto

type GetWeatherRequest struct{}

type GetWeatherResponse struct {
	City    string  `json:"city"`
	Weather Weather `json:"weather"`
}
