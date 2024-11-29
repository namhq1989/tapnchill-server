package dto

type SubscriptionPlan struct {
	PeriodText          string  `json:"periodText"`
	PriceID             string  `json:"priceId"`
	Amount              float64 `json:"amount"`
	DiscountID          string  `json:"discountId"`
	AfterDiscountAmount float64 `json:"afterDiscountAmount"`
	Token               string  `json:"token"`
}

type GetSubscriptionPlansRequest struct{}

type GetSubscriptionPlansResponse struct {
	IsEnabled bool               `json:"isEnabled"`
	Plans     []SubscriptionPlan `json:"plans"`
}
