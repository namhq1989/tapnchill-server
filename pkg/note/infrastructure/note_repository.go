package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoteRepository struct {
	db             *database.Database
	collectionName string
}

func NewNoteRepository(db *database.Database) NoteRepository {
	r := NoteRepository{
		db:             db,
		collectionName: database.Collections.Note,
	}
	r.ensureIndexes()
	return r
}

func (r NoteRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys: bson.D{{Key: "userId", Value: 1}, {Key: "createdAt", Value: 1}},
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r NoteRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r NoteRepository) Create(ctx *appcontext.AppContext, note domain.Note) error {
	doc, err := dbmodel.Note{}.FromDomain(note)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r NoteRepository) Update(ctx *appcontext.AppContext, note domain.Note) error {
	doc, err := dbmodel.Note{}.FromDomain(note)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r NoteRepository) Delete(ctx *appcontext.AppContext, noteID string) error {
	nid, err := database.ObjectIDFromString(noteID)
	if err != nil {
		return apperrors.Common.InvalidNote
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{"_id": nid})
	return err
}

func (r NoteRepository) CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return 0, apperrors.User.InvalidUserID
	}

	return r.collection().CountDocuments(ctx.Context(), bson.M{"userId": uid})
}

func (r NoteRepository) FindByID(ctx *appcontext.AppContext, noteID string) (*domain.Note, error) {
	nid, err := database.ObjectIDFromString(noteID)
	if err != nil {
		return nil, apperrors.Common.InvalidNote
	}

	var doc dbmodel.Note
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": nid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r NoteRepository) Sync(ctx *appcontext.AppContext, userID string, updatedAt time.Time, numOfNotes int64) ([]domain.Note, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var (
		condition = bson.M{
			"userId": uid,
			"updatedAt": bson.M{
				"$gt": updatedAt,
			},
		}
		result = make([]domain.Note, 0)
	)

	opts := options.Find().SetSort(bson.D{
		{Key: "createdAt", Value: 1},
	}).SetLimit(numOfNotes)
	cursor, err := r.collection().Find(ctx.Context(), condition, opts)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.Note
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
