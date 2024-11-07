package dto

type GetWeatherRequest struct{}

type GetWeatherResponse struct {
	Weather Weather `json:"weather"`
}
