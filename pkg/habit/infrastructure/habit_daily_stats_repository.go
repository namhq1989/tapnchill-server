package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HabitDailyStatsRepository struct {
	db             *database.Database
	collectionName string
}

func NewHabitDailyStatsRepository(db *database.Database) HabitDailyStatsRepository {
	r := HabitDailyStatsRepository{
		db:             db,
		collectionName: database.Collections.HabitDailyStats,
	}
	r.ensureIndexes()
	return r
}

func (r HabitDailyStatsRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "userId", Value: 1}, {Key: "date", Value: -1}},
				Options: options.Index().SetUnique(true),
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r HabitDailyStatsRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r HabitDailyStatsRepository) Create(ctx *appcontext.AppContext, stats domain.HabitDailyStats) error {
	doc, err := dbmodel.HabitDailyStats{}.FromDomain(stats)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r HabitDailyStatsRepository) Update(ctx *appcontext.AppContext, stats domain.HabitDailyStats) error {
	doc, err := dbmodel.HabitDailyStats{}.FromDomain(stats)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r HabitDailyStatsRepository) FindByID(ctx *appcontext.AppContext, statsID string) (*domain.HabitDailyStats, error) {
	sid, err := database.ObjectIDFromString(statsID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	var doc dbmodel.HabitDailyStats
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": sid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r HabitDailyStatsRepository) FindByDate(ctx *appcontext.AppContext, userID string, date time.Time) (*domain.HabitDailyStats, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.HabitDailyStats
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"userId": uid,
		"date":   date,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r HabitDailyStatsRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.HabitDailyStatsFilter) ([]domain.HabitDailyStats, error) {
	var (
		condition = bson.M{
			"userId": filter.UserID,
			"date": bson.M{
				"$gte": filter.FromDate,
			},
		}
		result = make([]domain.HabitDailyStats, 0)
	)

	cursor, err := r.collection().Find(ctx.Context(), condition, &options.FindOptions{
		Sort: bson.M{"date": -1},
	})
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.HabitDailyStats
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
