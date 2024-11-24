package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubscriptionHistoryRepository struct {
	db             *database.Database
	collectionName string
}

func NewSubscriptionHistoryRepository(db *database.Database) SubscriptionHistoryRepository {
	r := SubscriptionHistoryRepository{
		db:             db,
		collectionName: database.Collections.SubscriptionHistory,
	}
	r.ensureIndexes()
	return r
}

func (r SubscriptionHistoryRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "createdAt", Value: -1}},
			},
			{
				Keys:    bson.D{{Key: "sourceId", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r SubscriptionHistoryRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r SubscriptionHistoryRepository) Create(ctx *appcontext.AppContext, history domain.SubscriptionHistory) error {
	doc, err := dbmodel.SubscriptionHistory{}.FromDomain(history)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r SubscriptionHistoryRepository) Update(ctx *appcontext.AppContext, history domain.SubscriptionHistory) error {
	doc, err := dbmodel.SubscriptionHistory{}.FromDomain(history)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r SubscriptionHistoryRepository) FindBySourceID(ctx *appcontext.AppContext, sourceID string) (*domain.SubscriptionHistory, error) {
	var doc dbmodel.SubscriptionHistory
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"sourceId": sourceID,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}
