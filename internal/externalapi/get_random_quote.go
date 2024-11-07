package externalapi

import (
	"strconv"

	"github.com/namhq1989/go-utilities/appcontext"
)

type GetRandomQuoteResult struct {
	OriginalID string
	Content    string
	Author     string
}

type getRandomQuoteApiResult struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func (ea ExternalApi) GetRandomQuote(ctx *appcontext.AppContext) (*GetRandomQuoteResult, error) {
	var apiResult = &getRandomQuoteApiResult{}

	_, err := ea.quote.R().
		SetResult(&apiResult).
		Get("/api/quotes/random")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when get random quote", err, appcontext.Fields{})
		return nil, err
	}

	return &GetRandomQuoteResult{
		OriginalID: strconv.Itoa(apiResult.ID),
		Content:    apiResult.Quote,
		Author:     apiResult.Author,
	}, nil
}
