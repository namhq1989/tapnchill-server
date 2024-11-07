package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GoalRepository struct {
	db             *database.Database
	collectionName string
}

func NewGoalRepository(db *database.Database) GoalRepository {
	r := GoalRepository{
		db:             db,
		collectionName: database.Collections.Goal,
	}
	r.ensureIndexes()
	return r
}

func (r GoalRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "searchString", Value: "text"}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r GoalRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r GoalRepository) Create(ctx *appcontext.AppContext, goal domain.Goal) error {
	doc, err := dbmodel.Goal{}.FromDomain(goal)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r GoalRepository) Update(ctx *appcontext.AppContext, goal domain.Goal) error {
	doc, err := dbmodel.Goal{}.FromDomain(goal)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r GoalRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.GoalFilter) ([]domain.Goal, error) {
	var (
		condition = bson.M{
			"userId": filter.UserID,
			"createdAt": bson.M{
				"$lt": filter.Timestamp,
			},
		}
		result = make([]domain.Goal, 0)
	)

	if filter.Keyword != "" {
		condition["searchString"] = bson.M{"$text": bson.M{"$search": filter.Keyword}}
	}

	cursor, err := r.collection().Find(ctx.Context(), condition, nil)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.Goal
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}