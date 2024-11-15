package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HabitCompletionRepository struct {
	db             *database.Database
	collectionName string
}

func NewHabitCompletionRepository(db *database.Database) HabitCompletionRepository {
	r := HabitCompletionRepository{
		db:             db,
		collectionName: database.Collections.HabitCompletion,
	}
	r.ensureIndexes()
	return r
}

func (r HabitCompletionRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "habitId", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r HabitCompletionRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r HabitCompletionRepository) Create(ctx *appcontext.AppContext, completion domain.HabitCompletion) error {
	doc, err := dbmodel.HabitCompletion{}.FromDomain(completion)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}
