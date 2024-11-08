package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Goal struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserID       primitive.ObjectID `bson:"userId"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	SearchString string             `bson:"searchString"`
	Stats        GoalStats          `bson:"stats"`
	IsCompleted  bool               `bson:"isCompleted"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

type GoalStats struct {
	TotalTask          int `bson:"totalTask"`
	TotalCompletedTask int `bson:"totalCompletedTask"`
}

func (g Goal) ToDomain() domain.Goal {
	return domain.Goal{
		ID:           g.ID.Hex(),
		UserID:       g.UserID.Hex(),
		Name:         g.Name,
		Description:  g.Description,
		SearchString: g.SearchString,
		Stats: domain.GoalStats{
			TotalTask:          g.Stats.TotalTask,
			TotalCompletedTask: g.Stats.TotalCompletedTask,
		},
		IsCompleted: g.IsCompleted,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}
}

func (Goal) FromDomain(goal domain.Goal) (*Goal, error) {
	id, err := database.ObjectIDFromString(goal.ID)
	if err != nil {
		return nil, apperrors.Task.InvalidGoalID
	}

	uid, err := database.ObjectIDFromString(goal.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &Goal{
		ID:           id,
		UserID:       uid,
		Name:         goal.Name,
		Description:  goal.Description,
		SearchString: goal.SearchString,
		Stats: GoalStats{
			TotalTask:          goal.Stats.TotalTask,
			TotalCompletedTask: goal.Stats.TotalCompletedTask,
		},
		IsCompleted: goal.IsCompleted,
		CreatedAt:   goal.CreatedAt,
		UpdatedAt:   goal.UpdatedAt,
	}, nil
}
