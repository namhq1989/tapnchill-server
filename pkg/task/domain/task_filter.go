package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskFilter struct {
	UserID    primitive.ObjectID
	GoalID    primitive.ObjectID
	Keyword   string
	Timestamp time.Time
	Limit     int64
}

func NewTaskFilter(userID, keyword, pt string) (*TaskFilter, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	gid, err := database.ObjectIDFromString(keyword)
	if err != nil {
		return nil, apperrors.Common.InvalidGoal
	}

	pageToken := pagetoken.Decode(pt)

	return &TaskFilter{
		UserID:    uid,
		GoalID:    gid,
		Keyword:   keyword,
		Timestamp: pageToken.Timestamp,
		Limit:     20,
	}, nil
}
