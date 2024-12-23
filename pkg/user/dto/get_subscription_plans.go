package dto

type SubscriptionPlan struct {
	ID                  string  `json:"id"`
	Amount              float64 `json:"amount"`
	AfterDiscountAmount float64 `json:"afterDiscountAmount"`
}

type PlanLimitation struct {
	Free int64 `json:"free"`
	Pro  int64 `json:"pro"`
}

type ResourcesLimitation struct {
	Goal   PlanLimitation `json:"goal"`
	Task   PlanLimitation `json:"task"`
	Habit  PlanLimitation `json:"habit"`
	Note   PlanLimitation `json:"note"`
	QRCode PlanLimitation `json:"qrCode"`
}

type GetSubscriptionPlansRequest struct{}

type GetSubscriptionPlansResponse struct {
	IsEnabled           bool                `json:"isEnabled"`
	Plans               []SubscriptionPlan  `json:"plans"`
	ResourcesLimitation ResourcesLimitation `json:"resourcesLimitation"`
}
