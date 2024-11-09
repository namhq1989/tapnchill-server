package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Plan          string             `bson:"plan"`
	AuthProviders []AuthProvider     `bson:"authProviders"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}

func (u User) ToDomain() domain.User {
	user := domain.User{
		ID:            u.ID.Hex(),
		Name:          u.Name,
		Plan:          domain.ToPlan(u.Plan),
		AuthProviders: make([]domain.AuthProvider, 0),
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}

	for _, ap := range u.AuthProviders {
		user.AuthProviders = append(user.AuthProviders, ap.ToDomain())
	}

	return user
}

func (User) FromDomain(user domain.User) (*User, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, err
	}

	authProviders := make([]AuthProvider, 0)
	for _, ap := range user.AuthProviders {
		authProviders = append(authProviders, AuthProvider{}.FromDomain(ap))
	}

	return &User{
		ID:            id,
		Name:          user.Name,
		Plan:          user.Plan.String(),
		AuthProviders: authProviders,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}, nil
}
