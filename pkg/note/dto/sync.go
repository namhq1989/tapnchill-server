package dto

type SyncRequest struct {
	LastUpdatedAt string `query:"lastUpdatedAt"`
}

type SyncResponse struct {
	Notes []Note `json:"notes"`
	Limit int64  `json:"limit"`
}
