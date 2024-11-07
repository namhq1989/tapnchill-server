package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/infrastructure/dbmodel"

	"github.com/namhq1989/tapnchill-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuoteRepository struct {
	db             *database.Database
	collectionName string
}

func NewQuoteRepository(db *database.Database) QuoteRepository {
	r := QuoteRepository{
		db:             db,
		collectionName: database.Collections.Quote,
	}
	r.ensureIndexes()
	return r
}

func (r QuoteRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "createdAt", Value: -1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r QuoteRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r QuoteRepository) Create(ctx *appcontext.AppContext, quote domain.Quote) error {
	doc, err := dbmodel.Quote{}.FromDomain(quote)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r QuoteRepository) FindLatest(ctx *appcontext.AppContext) (*domain.Quote, error) {
	var doc dbmodel.Quote
	if err := r.collection().FindOne(ctx.Context(), bson.M{}, &options.FindOneOptions{
		Sort: bson.M{"createdAt": -1},
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}
