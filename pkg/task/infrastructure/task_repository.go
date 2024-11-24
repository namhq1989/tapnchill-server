package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	db             *database.Database
	collectionName string
}

func NewTaskRepository(db *database.Database) TaskRepository {
	r := TaskRepository{
		db:             db,
		collectionName: database.Collections.Task,
	}
	r.ensureIndexes()
	return r
}

func (r TaskRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "userId", Value: 1},
					{Key: "goalId", Value: 1},
					{Key: "searchString", Value: "text"},
					{Key: "createdAt", Value: -1},
				},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r TaskRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r TaskRepository) Create(ctx *appcontext.AppContext, task domain.Task) error {
	doc, err := dbmodel.Task{}.FromDomain(task)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r TaskRepository) Update(ctx *appcontext.AppContext, task domain.Task) error {
	doc, err := dbmodel.Task{}.FromDomain(task)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r TaskRepository) Delete(ctx *appcontext.AppContext, taskID string) error {
	tid, err := database.ObjectIDFromString(taskID)
	if err != nil {
		return apperrors.Task.InvalidGoalID
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{"_id": tid})
	return err
}

func (r TaskRepository) CountByGoalID(ctx *appcontext.AppContext, goalID string) (int64, error) {
	gid, err := database.ObjectIDFromString(goalID)
	if err != nil {
		return 0, apperrors.Task.InvalidGoalID
	}

	return r.collection().CountDocuments(ctx.Context(), bson.M{"goalId": gid})
}

func (r TaskRepository) FindByID(ctx *appcontext.AppContext, taskID string) (*domain.Task, error) {
	tid, err := database.ObjectIDFromString(taskID)
	if err != nil {
		return nil, apperrors.Task.InvalidTaskID
	}

	var doc dbmodel.Task
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": tid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r TaskRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.TaskFilter) ([]domain.Task, error) {
	var (
		condition = bson.M{
			"userId": filter.UserID,
			"createdAt": bson.M{
				"$lt": filter.Timestamp,
			},
		}
		result = make([]domain.Task, 0)
	)

	if !filter.GoalID.IsZero() {
		condition["goalId"] = filter.GoalID
	}

	if filter.Status.IsValid() {
		condition["status"] = filter.Status
	}

	if filter.Keyword != "" {
		condition["searchString"] = bson.M{"$text": bson.M{"$search": filter.Keyword}}
	}

	cursor, err := r.collection().Find(ctx.Context(), condition, &options.FindOptions{
		Sort: bson.M{"createdAt": -1},
	})
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.Task
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
