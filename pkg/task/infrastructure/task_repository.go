package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/infrastructure/dbmodel"

	"github.com/namhq1989/tapnchill-server/internal/database"
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

	if filter.Keyword != "" {
		condition["searchString"] = bson.M{"$text": bson.M{"$search": filter.Keyword}}
	}

	cursor, err := r.collection().Find(ctx.Context(), condition, nil)
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
