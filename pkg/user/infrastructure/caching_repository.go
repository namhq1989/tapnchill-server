package infrastructure

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/caching"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

type CachingRepository struct {
	caching caching.Operations

	domain                                string
	userDuration                          time.Duration
	lemonsqueezyCustomerPortalURLDuration time.Duration
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching:                               caching,
		domain:                                "user",
		userDuration:                          12 * time.Hour,
		lemonsqueezyCustomerPortalURLDuration: 23 * time.Hour,
	}
}

// USER

func (r CachingRepository) GetUserByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	key := r.generateUserByIDKey(userID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result *domain.User
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r CachingRepository) SetUserByID(ctx *appcontext.AppContext, userID string, user domain.User) error {
	key := r.generateUserByIDKey(userID)
	r.caching.SetTTL(ctx, key, user, r.userDuration)
	return nil
}

func (r CachingRepository) DeleteUserByID(ctx *appcontext.AppContext, userID string) error {
	key := r.generateUserByIDKey(userID)
	_, _ = r.caching.Del(ctx, key)
	return nil
}

func (r CachingRepository) generateUserByIDKey(userID string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("user:%s", userID))
}

// LEMONSQUEEZY CUSTOMER PORTAL URL

func (r CachingRepository) GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, customerID string) (*string, error) {
	key := r.generateLemonsqueezyCustomerPortalURLKey(customerID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result *string
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r CachingRepository) SetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, customerID, url string) error {
	key := r.generateLemonsqueezyCustomerPortalURLKey(customerID)
	r.caching.SetTTL(ctx, key, url, r.lemonsqueezyCustomerPortalURLDuration)
	return nil
}

func (r CachingRepository) generateLemonsqueezyCustomerPortalURLKey(customer string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("payment:lemonsqueezy:%s:portal_url", customer))
}
