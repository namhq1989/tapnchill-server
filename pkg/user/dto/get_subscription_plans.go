package dto

type SubscriptionPlan struct {
	ID                  string  `json:"id"`
	Amount              float64 `json:"amount"`
	AfterDiscountAmount float64 `json:"afterDiscountAmount"`
}

type GetSubscriptionPlansRequest struct{}

type GetSubscriptionPlansResponse struct {
	IsEnabled bool               `json:"isEnabled"`
	Plans     []SubscriptionPlan `json:"plans"`
}
