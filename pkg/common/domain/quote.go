package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

type QuoteRepository interface {
	Create(ctx *appcontext.AppContext, quote Quote) error
	FindLatest(ctx *appcontext.AppContext) (*Quote, error)
}

type Quote struct {
	ID         string
	OriginalID string
	Content    string
	Author     string
	CreatedAt  time.Time
}

func NewQuote(originalID, content, author string) (*Quote, error) {
	if originalID == "" || content == "" || author == "" {
		return nil, apperrors.Common.InvalidFeedback
	}

	return &Quote{
		ID:         database.NewStringID(),
		OriginalID: originalID,
		Content:    content,
		Author:     author,
		CreatedAt:  time.Now(),
	}, nil
}
