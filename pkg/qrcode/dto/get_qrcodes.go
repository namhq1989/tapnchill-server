package dto

type GetQRCodesRequest struct {
	PageToken string `query:"pageToken"`
}

type GetQRCodesResponse struct {
	QRCodes       []QRCode `json:"qrCodes"`
	NextPageToken string   `json:"nextPageToken"`
}
