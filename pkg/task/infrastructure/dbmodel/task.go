package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserID       primitive.ObjectID `bson:"userId"`
	GoalID       primitive.ObjectID `bson:"goalId"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	SearchString string             `bson:"searchString"`
	DueDate      *time.Time         `bson:"dueDate"`
	IsCompleted  bool               `bson:"isCompleted"`
	CreatedAt    time.Time          `bson:"createdAt"`
	CompletedAt  *time.Time         `bson:"completedAt"`
}

func (t Task) ToDomain() domain.Task {
	return domain.Task{
		ID:           t.ID.Hex(),
		UserID:       t.UserID.Hex(),
		GoalID:       t.GoalID.Hex(),
		Name:         t.Name,
		Description:  t.Description,
		SearchString: t.SearchString,
		DueDate:      t.DueDate,
		IsCompleted:  t.IsCompleted,
		CreatedAt:    t.CreatedAt,
		CompletedAt:  t.CompletedAt,
	}
}

func (Task) FromDomain(task domain.Task) (*Task, error) {
	id, err := database.ObjectIDFromString(task.ID)
	if err != nil {
		return nil, apperrors.Task.InvalidTaskID
	}

	uid, err := database.ObjectIDFromString(task.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	gid, err := database.ObjectIDFromString(task.GoalID)
	if err != nil {
		return nil, apperrors.Task.InvalidGoalID
	}

	return &Task{
		ID:           id,
		UserID:       uid,
		GoalID:       gid,
		Name:         task.Name,
		Description:  task.Description,
		SearchString: task.SearchString,
		DueDate:      task.DueDate,
		IsCompleted:  task.IsCompleted,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
	}, nil
}
