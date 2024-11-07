package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoalFilter struct {
	UserID    primitive.ObjectID
	Keyword   string
	Timestamp time.Time
	Limit     int64
}

func NewGoalFilter(userID, keyword, pt string) (*GoalFilter, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	pageToken := pagetoken.Decode(pt)

	return &GoalFilter{
		UserID:    uid,
		Keyword:   keyword,
		Timestamp: pageToken.Timestamp,
		Limit:     10,
	}, nil
}
