package dto

type GetMeRequest struct{}

type GetMeResponse struct {
	Ip           string       `json:"ip"`
	Subscription Subscription `json:"subscription"`
}
