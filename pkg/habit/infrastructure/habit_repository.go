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

type HabitRepository struct {
	db             *database.Database
	collectionName string
}

func NewHabitRepository(db *database.Database) HabitRepository {
	r := HabitRepository{
		db:             db,
		collectionName: database.Collections.Habit,
	}
	r.ensureIndexes()
	return r
}

func (r HabitRepository) ensureIndexes() {
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

func (r HabitRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r HabitRepository) Create(ctx *appcontext.AppContext, habit domain.Habit) error {
	doc, err := dbmodel.Habit{}.FromDomain(habit)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r HabitRepository) Update(ctx *appcontext.AppContext, habit domain.Habit) error {
	doc, err := dbmodel.Habit{}.FromDomain(habit)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r HabitRepository) Delete(ctx *appcontext.AppContext, habitID string) error {
	hid, err := database.ObjectIDFromString(habitID)
	if err != nil {
		return apperrors.Habit.InvalidID
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{"_id": hid})
	return err
}

func (r HabitRepository) FindByID(ctx *appcontext.AppContext, habitID string) (*domain.Habit, error) {
	hid, err := database.ObjectIDFromString(habitID)
	if err != nil {
		return nil, apperrors.Habit.InvalidID
	}

	var doc dbmodel.Habit
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": hid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r HabitRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.HabitFilter) ([]domain.Habit, error) {
	var (
		condition = bson.M{
			"userId": filter.UserID,
		}
		result = make([]domain.Habit, 0)
	)

	opts := options.Find().SetSort(bson.D{
		{Key: "sortOrder", Value: 1},
		{Key: "createdAt", Value: -1},
	})
	cursor, err := r.collection().Find(ctx.Context(), condition, opts)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.Habit
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
