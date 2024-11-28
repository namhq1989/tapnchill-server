package dto

type FastspringRequest struct {
	Events []FastspringEvent `json:"events"`
}

type FastspringEvent struct {
	ID   string              `json:"id"`
	Type string              `json:"type"`
	Data FastspringEventData `json:"data"`
}

type FastspringEventData struct {
	Account        string                  `json:"account"`
	Subscription   string                  `json:"subscription"`
	Product        string                  `json:"product"`
	NextChargeDate string                  `json:"nextChargeDateDisplayISO8601"`
	Tags           FastspringEventDataTags `json:"tags"`
}

type FastspringEventDataTags struct {
	UserID string `json:"userId"`
}

type FastspringResponse struct{}
