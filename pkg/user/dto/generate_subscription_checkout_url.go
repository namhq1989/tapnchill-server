package dto

type GenerateSubscriptionCheckoutURLRequest struct {
	SubscriptionID string `json:"subscriptionId"`
}

type GenerateSubscriptionCheckoutURLResponse struct {
	CheckoutURL string `json:"checkoutUrl"`
}
