package infrastructure

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/infrastructure/dbmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db             *database.Database
	collectionName string
	checksumSecret []byte
}

func NewUserRepository(db *database.Database, checksumSecret string) UserRepository {
	r := UserRepository{
		db:             db,
		collectionName: database.Collections.User,
		checksumSecret: []byte(checksumSecret),
	}
	r.ensureIndexes()
	return r
}

func (r UserRepository) ensureIndexes() {
	var (
		ctx     = context.Background()
		opts    = options.CreateIndexes().SetMaxTime(time.Minute * 30)
		indexes = []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "authProviders.provider", Value: 1}, {Key: "authProviders.id", Value: 1}},
				Options: options.Index().SetUnique(true).SetSparse(true),
			},
			{
				Keys:    bson.D{{Key: "authProviders.email", Value: 1}},
				Options: options.Index().SetSparse(true),
			},
		}
	)

	if _, err := r.collection().Indexes().CreateMany(ctx, indexes, opts); err != nil {
		fmt.Printf("index collection %s err: %v \n", r.collectionName, err)
	}
}

func (r UserRepository) collection() *mongo.Collection {
	return r.db.GetCollection(r.collectionName)
}

func (r UserRepository) Create(ctx *appcontext.AppContext, user domain.User) error {
	doc, err := dbmodel.User{}.FromDomain(user)
	if err != nil {
		return err
	}

	_, err = r.collection().InsertOne(ctx.Context(), doc)
	return err
}

func (r UserRepository) Update(ctx *appcontext.AppContext, user domain.User) error {
	doc, err := dbmodel.User{}.FromDomain(user)
	if err != nil {
		return err
	}

	_, err = r.collection().UpdateByID(ctx.Context(), doc.ID, bson.M{"$set": doc})
	return err
}

func (r UserRepository) Delete(ctx *appcontext.AppContext, userID string) error {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return apperrors.User.InvalidUserID
	}

	_, err = r.collection().DeleteOne(ctx.Context(), bson.M{"_id": uid})
	return err
}

func (r UserRepository) FindByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	var doc dbmodel.User
	if err = r.collection().FindOne(ctx.Context(), bson.M{
		"_id": uid,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserRepository) FindByAuthProviderID(ctx *appcontext.AppContext, id string) (*domain.User, error) {
	var doc dbmodel.User
	if err := r.collection().FindOne(ctx.Context(), bson.M{
		"authProviders.id": id,
	}).Decode(&doc); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	result := doc.ToDomain()
	return &result, nil
}

func (r UserRepository) ValidateAnonymousChecksum(_ *appcontext.AppContext, clientID, checksum string) bool {
	mac := hmac.New(sha256.New, r.checksumSecret)
	mac.Write([]byte(clientID))
	expectedMAC := mac.Sum(nil)

	providedMAC, err := hex.DecodeString(checksum)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedMAC, providedMAC)
}
