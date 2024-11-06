package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeedbackRepository struct {
	db             *database.Database
	collectionName string
}

func NewFeedbackRepository(db *database.Database) FeedbackRepository {
	r := FeedbackRepository{
		db:             db,
		collectionName: database.Collections.Feedback,
	}
	r.ensureIndexes()
	return r
}

func (r FeedbackRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r FeedbackRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r FeedbackRepository) Create(ctx *appcontext.AppContext, feedback domain.Feedback) error {
	doc, err := dbmodel.Feedback{}.FromDomain(feedback)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}
