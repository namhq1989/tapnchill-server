package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID  `bson:"_id"`
	BelongsTo     *primitive.ObjectID `bson:"belongsTo"`
	Name          string              `bson:"name"`
	Subscription  UserSubscription    `bson:"subscription"`
	AuthProviders []AuthProvider      `bson:"authProviders"`
	CreatedAt     time.Time           `bson:"createdAt"`
	UpdatedAt     time.Time           `bson:"updatedAt"`
}

func (u User) ToDomain() domain.User {
	user := domain.User{
		ID:            u.ID.Hex(),
		BelongsTo:     nil,
		Name:          u.Name,
		Subscription:  u.Subscription.ToDomain(),
		AuthProviders: make([]domain.AuthProvider, 0),
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}

	if u.BelongsTo != nil {
		hexID := u.BelongsTo.Hex()
		user.BelongsTo = &hexID
	}

	for _, ap := range u.AuthProviders {
		user.AuthProviders = append(user.AuthProviders, ap.ToDomain())
	}

	return user
}

func (User) FromDomain(user domain.User) (*User, error) {
	id, err := database.ObjectIDFromString(user.ID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	authProviders := make([]AuthProvider, 0)
	for _, ap := range user.AuthProviders {
		authProviders = append(authProviders, AuthProvider{}.FromDomain(ap))
	}

	var belongsTo primitive.ObjectID
	if user.BelongsTo != nil {
		bid, bErr := database.ObjectIDFromString(*user.BelongsTo)
		if bErr != nil {
			return nil, apperrors.User.InvalidUserID
		}

		belongsTo = bid
	}

	return &User{
		ID:            id,
		Name:          user.Name,
		BelongsTo:     &belongsTo,
		Subscription:  UserSubscription{}.FromDomain(user.Subscription),
		AuthProviders: authProviders,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}, nil
}
