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
	Status    TaskStatus
	Keyword   string
	Timestamp time.Time
	Limit     int64
}

func NewTaskFilter(userID, goalID, status, keyword, pt string) (*TaskFilter, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	pageToken := pagetoken.Decode(pt)

	filter := &TaskFilter{
		UserID:    uid,
		Keyword:   keyword,
		Timestamp: pageToken.Timestamp,
		Limit:     20,
	}

	if goalID != "" {
		gid, gErr := database.ObjectIDFromString(goalID)
		if gErr != nil {
			return nil, apperrors.Task.InvalidGoalID
		}
		filter.GoalID = gid
	}

	dStatus := ToTaskStatus(status)
	if dStatus.IsValid() {
		filter.Status = dStatus
	}

	return filter, nil
}
