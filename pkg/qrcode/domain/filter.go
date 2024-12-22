package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
)

type QRCodeFilter struct {
	UserID    string
	Timestamp time.Time
	Limit     int64
}

func NewQRCodeFilter(userID string, pt string) (*QRCodeFilter, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	pageToken := pagetoken.Decode(pt)
	return &QRCodeFilter{
		UserID:    userID,
		Timestamp: pageToken.Timestamp,
		Limit:     20,
	}, nil
}
