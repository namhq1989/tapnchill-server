package dto

type LemonsqueezyRequest struct {
	Data LemonsqueezyRequestData `json:"data"`
	Meta LemonsqueezyRequestMeta `json:"meta"`
}

type LemonsqueezyRequestData struct {
	InvoiceID string `json:"id"`
}

type LemonsqueezyRequestMeta struct {
	EventName  string                            `json:"event_name"`
	CustomData LemonsqueezyRequestMetaCustomData `json:"custom_data"`
}

type LemonsqueezyRequestMetaCustomData struct {
	UserID string `json:"user_id"`
}

type LemonsqueezyResponse struct{}
