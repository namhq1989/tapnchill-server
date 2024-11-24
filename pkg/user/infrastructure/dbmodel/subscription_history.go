package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionHistory struct {
	ID               primitive.ObjectID `bson:"_id"`
	UserID           primitive.ObjectID `bson:"userId"`
	SourceID         string             `bson:"sourceId"`
	SourceName       string             `bson:"sourceName"`
	SourceCustomerID string             `bson:"sourceCustomerId"`
	Items            []string           `bson:"items"`
	Status           string             `bson:"status"`
	CreatedAt        time.Time          `bson:"createdAt"`
	NextBilledAt     time.Time          `bson:"nextBilledAt"`
}

func (s SubscriptionHistory) ToDomain() domain.SubscriptionHistory {
	return domain.SubscriptionHistory{
		ID:               s.ID.Hex(),
		UserID:           s.UserID.Hex(),
		SourceID:         s.SourceID,
		SourceName:       s.SourceName,
		SourceCustomerID: s.SourceCustomerID,
		Items:            s.Items,
		Status:           domain.SubscriptionHistoryStatus(s.Status),
		CreatedAt:        s.CreatedAt,
		NextBilledAt:     s.NextBilledAt,
	}
}

func (SubscriptionHistory) FromDomain(history domain.SubscriptionHistory) (*SubscriptionHistory, error) {
	id, err := database.ObjectIDFromString(history.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(history.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &SubscriptionHistory{
		ID:               id,
		UserID:           uid,
		SourceID:         history.SourceID,
		SourceName:       history.SourceName,
		SourceCustomerID: history.SourceCustomerID,
		Items:            history.Items,
		Status:           history.Status.String(),
		CreatedAt:        history.CreatedAt,
		NextBilledAt:     history.NextBilledAt,
	}, nil
}
