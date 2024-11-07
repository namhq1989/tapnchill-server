package dto

type GetQuoteRequest struct{}

type GetQuoteResponse struct {
	Quote Quote `json:"quote"`
}
