package dto

type PaddleRequest struct {
	Data      PaddleRequestData `json:"data"`
	EventType string            `json:"event_type"`
}

type PaddleRequestData struct {
	ID            string                      `json:"id"`
	BillingPeriod PaddleBillingPeriod         `json:"billing_period"`
	CustomData    PaddleRequestDataCustomData `json:"custom_data"`
}

type PaddleRequestDataItem struct {
	PriceID string `json:"price_id"`
}

type PaddleBillingPeriod struct {
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
}

type PaddleRequestDataCustomData struct {
	UserID string `json:"userId"`
}

type PaddleResponse struct{}
