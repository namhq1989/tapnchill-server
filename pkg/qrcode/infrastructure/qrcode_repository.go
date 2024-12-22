package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QRCodeRepository struct {
	db             *database.Database
	collectionName string
}

func NewQRCodeRepository(db *database.Database) QRCodeRepository {
	r := QRCodeRepository{
		db:             db,
		collectionName: database.Collections.QRCode,
	}
	r.ensureIndexes()
	return r
}

func (r QRCodeRepository) ensureIndexes() {
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

func (r QRCodeRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r QRCodeRepository) Create(ctx *appcontext.AppContext, qrCode domain.QRCode) error {
	doc, err := dbmodel.QRCode{}.FromDomain(qrCode)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r QRCodeRepository) Update(ctx *appcontext.AppContext, qrCode domain.QRCode) error {
	doc, err := dbmodel.QRCode{}.FromDomain(qrCode)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r QRCodeRepository) Delete(ctx *appcontext.AppContext, qrCodeID string) error {
	hid, err := database.ObjectIDFromString(qrCodeID)
	if err != nil {
		return apperrors.Common.InvalidID
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{"_id": hid})
	return err
}

func (r QRCodeRepository) CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return 0, apperrors.User.InvalidUserID
	}

	return r.collection().CountDocuments(ctx.Context(), bson.M{"userId": uid})
}

func (r QRCodeRepository) FindByID(ctx *appcontext.AppContext, qrCodeID string) (*domain.QRCode, error) {
	hid, err := database.ObjectIDFromString(qrCodeID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	var doc dbmodel.QRCode
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

func (r QRCodeRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.QRCodeFilter) ([]domain.QRCode, error) {
	uid, err := database.ObjectIDFromString(filter.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var (
		condition = bson.M{
			"userId": uid,
		}
		result = make([]domain.QRCode, 0)
	)

	if !filter.Timestamp.IsZero() {
		condition["createdAt"] = bson.M{
			"$lt": filter.Timestamp,
		}
	}

	opts := options.Find().SetSort(bson.D{
		{Key: "createdAt", Value: -1},
	}).SetLimit(filter.Limit)
	cursor, err := r.collection().Find(ctx.Context(), condition, opts)
	if err != nil {
		return result, err
	}
	defer func() { _ = cursor.Close(ctx.Context()) }()

	var docs []dbmodel.QRCode
	if err = cursor.All(ctx.Context(), &docs); err != nil {
		return result, err
	}

	for _, doc := range docs {
		result = append(result, doc.ToDomain())
	}
	return result, nil
}
