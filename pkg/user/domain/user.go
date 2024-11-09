package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type UserRepository interface {
	Create(ctx *appcontext.AppContext, user User) error
	Update(ctx *appcontext.AppContext, user User) error
	Delete(ctx *appcontext.AppContext, userID string) error
	FindByID(ctx *appcontext.AppContext, userID string) (*User, error)
	ValidateAnonymousChecksum(_ *appcontext.AppContext, clientID, checksum string) bool
}

type User struct {
	ID            string
	Name          string
	Plan          Plan
	AuthProviders []AuthProvider
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(clientID, source string) (*User, error) {
	if clientID == "" {
		return nil, apperrors.User.InvalidClientID
	}

	authProviders := make([]AuthProvider, 0)
	if source == "" || source == AuthProviderExtension {
		source = AuthProviderExtension
		authProviders = append(authProviders, AuthProvider{
			Provider: source,
			ID:       clientID,
			Name:     clientID,
			Email:    "",
		})
	}

	return &User{
		ID:            database.NewStringID(),
		Name:          clientID,
		Plan:          PlanFree,
		AuthProviders: authProviders,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func (u *User) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	u.Name = name
	u.SetUpdatedAt()
	return nil
}

func (u *User) SetPlan(plan Plan) {
	u.Plan = plan
	u.SetUpdatedAt()
}

func (u *User) AddAuthProvider(provider AuthProvider) {
	// check provider existence by provider.provider
	for _, p := range u.AuthProviders {
		if p.Provider == provider.Provider {
			return
		}
	}

	u.AuthProviders = append(u.AuthProviders, provider)
	u.SetUpdatedAt()
}

func (u *User) SetUpdatedAt() {
	u.UpdatedAt = manipulation.NowUTC()
}
