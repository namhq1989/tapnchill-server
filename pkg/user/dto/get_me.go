package dto

type GetMeRequest struct{}

type GetMeResponse struct {
	Ip           string           `json:"ip"`
	Subscription UserSubscription `json:"subscription"`
}
