package dto

type PaddleRequest struct {
	Data      PaddleRequestData `json:"data"`
	EventType string            `json:"event_type"`
}

type PaddleRequestData struct {
	ID             string                      `json:"id"`
	Items          []PaddleRequestDataItem     `json:"items"`
	CustomData     PaddleRequestDataCustomData `json:"custom_data"`
	NextBilledAt   string                      `json:"next_billed_at,omitempty"`
	SubscriptionID string                      `json:"subscription_id,omitempty"`
	CustomerID     string                      `json:"customer_id"`
}

type PaddleRequestDataItem struct {
	Price PaddleRequestDataItemPrice `json:"price"`
}

type PaddleRequestDataItemPrice struct {
	ID string `json:"id"`
}

type PaddleRequestDataCustomData struct {
	UserID string `json:"userId"`
}

type PaddleResponse struct{}
