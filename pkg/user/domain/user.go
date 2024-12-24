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
	FindByAuthProviderID(ctx *appcontext.AppContext, id string) (*User, error)
	ValidateAnonymousChecksum(ctx *appcontext.AppContext, clientID, checksum string) bool
	DowngradeAllExpiredSubscriptions(ctx *appcontext.AppContext) (int64, error)
}

type User struct {
	ID            string
	BelongsTo     *string
	Name          string
	Subscription  UserSubscription
	AuthProviders []AuthProvider
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewExtensionUser(clientID string) (*User, error) {
	if clientID == "" {
		return nil, apperrors.User.InvalidClientID
	}

	return &User{
		ID:        database.NewStringID(),
		BelongsTo: nil,
		Name:      clientID,
		Subscription: UserSubscription{
			Plan:       PlanFree,
			Expiry:     nil,
			CustomerID: "",
		},
		AuthProviders: []AuthProvider{
			{
				Provider: AuthProviderExtension,
				ID:       clientID,
				Name:     clientID,
				Email:    "",
			},
		},
		CreatedAt: manipulation.NowUTC(),
		UpdatedAt: manipulation.NowUTC(),
	}, nil
}

func NewGoogleUser(id, email, name string) (*User, error) {
	if id == "" {
		return nil, apperrors.Auth.InvalidGoogleToken
	}

	if name == "" {
		name = id
	}

	return &User{
		ID:   database.NewStringID(),
		Name: name,
		Subscription: UserSubscription{
			Plan:       PlanFree,
			Expiry:     nil,
			CustomerID: "",
		},
		AuthProviders: []AuthProvider{
			{
				Provider: AuthProviderGoogle,
				ID:       id,
				Name:     name,
				Email:    email,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) SetBelongsTo(userID string) error {
	if !database.IsValidObjectID(userID) {
		return apperrors.User.InvalidUserID
	}

	u.BelongsTo = &userID
	u.UpdatedAt = manipulation.NowUTC()
	return nil
}

func (u *User) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	u.Name = name
	u.SetUpdatedAt()
	return nil
}

func (u *User) IsFreePlan() bool {
	return u.Subscription.Plan == PlanFree
}

func (u *User) IsProPlan() bool {
	return u.Subscription.Plan == PlanPro
}

func (u *User) SetPlanFree() {
	u.Subscription.Plan = PlanFree
	u.Subscription.Expiry = nil
	u.SetUpdatedAt()
}

func (u *User) SetPlanPro(expiry time.Time) {
	u.Subscription.Plan = PlanPro
	u.Subscription.Expiry = &expiry
	u.SetUpdatedAt()
}

func (u *User) SetSubscriptionCustomerID(customerID string) {
	u.Subscription.CustomerID = customerID
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
