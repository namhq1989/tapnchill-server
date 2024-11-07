package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quote struct {
	ID         primitive.ObjectID `bson:"_id"`
	OriginalID string             `bson:"originalId"`
	Content    string             `bson:"content"`
	Author     string             `bson:"author"`
	CreatedAt  time.Time          `bson:"createdAt"`
}

func (q Quote) ToDomain() domain.Quote {
	return domain.Quote{
		ID:         q.ID.Hex(),
		OriginalID: q.OriginalID,
		Content:    q.Content,
		Author:     q.Author,
		CreatedAt:  q.CreatedAt,
	}
}

func (Quote) FromDomain(quote domain.Quote) (*Quote, error) {
	id, err := database.ObjectIDFromString(quote.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidQuote
	}

	return &Quote{
		ID:         id,
		OriginalID: quote.OriginalID,
		Content:    quote.Content,
		Author:     quote.Author,
		CreatedAt:  quote.CreatedAt,
	}, nil
}
