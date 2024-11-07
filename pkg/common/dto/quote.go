package dto

import "github.com/namhq1989/tapnchill-server/pkg/common/domain"

type Quote struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (Quote) FromDomain(quote domain.Quote) Quote {
	return Quote{
		ID:      quote.ID,
		Content: quote.Content,
		Author:  quote.Author,
	}
}
