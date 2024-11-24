package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type SubscriptionHistoryRepository interface {
	Create(ctx *appcontext.AppContext, history SubscriptionHistory) error
	Update(ctx *appcontext.AppContext, history SubscriptionHistory) error
	FindBySourceID(ctx *appcontext.AppContext, sourceID string) (*SubscriptionHistory, error)
}

type SubscriptionHistoryStatus string

const (
	SubscriptionHistoryStatusActivated SubscriptionHistoryStatus = "activated"
	SubscriptionHistoryStatusPending   SubscriptionHistoryStatus = "pending"
)

func (s SubscriptionHistoryStatus) String() string {
	return string(s)
}

type SubscriptionHistory struct {
	ID               string
	UserID           string
	SourceID         string
	SourceName       string
	SourceCustomerID string
	Items            []string
	Status           SubscriptionHistoryStatus
	CreatedAt        time.Time
	NextBilledAt     time.Time
}

func NewSubscriptionHistory(userID string, sourceID string, sourceName string, sourceCustomerID string, items []string, nextBilledAt time.Time) (*SubscriptionHistory, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if sourceID == "" || len(items) == 0 {
		return nil, apperrors.Common.BadRequest
	}

	return &SubscriptionHistory{
		ID:               database.NewStringID(),
		UserID:           userID,
		SourceID:         sourceID,
		SourceName:       sourceName,
		SourceCustomerID: sourceCustomerID,
		Items:            items,
		Status:           SubscriptionHistoryStatusPending,
		CreatedAt:        manipulation.NowUTC(),
		NextBilledAt:     nextBilledAt,
	}, nil
}

func (h *SubscriptionHistory) SetActivated() {
	h.Status = SubscriptionHistoryStatusActivated
}
