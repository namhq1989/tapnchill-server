package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
)

func (s Service) GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, userID string) (*string, error) {
	ctx.Logger().Info("find user by id", appcontext.Fields{"userID": userID})
	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		ctx.Logger().Error("failed to find user by id", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found, respond")
		return nil, apperrors.Common.NotFound
	}

	customerID := user.Subscription.CustomerID
	if customerID == "" {
		ctx.Logger().ErrorText("customer id not found, respond")
		return nil, nil
	}

	ctx.Logger().Text("find portal url in caching")
	url, err := s.cachingRepository.GetLemonsqueezyCustomerPortalURL(ctx, customerID)
	if url != nil {
		ctx.Logger().Text("found in caching, respond")
		return url, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find portal url in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find portal url with Lemon Squeezy api")
	url, err = s.externalAPIRepository.GetLemonsqueezyCustomerPortalURL(ctx, customerID)
	if err != nil {
		ctx.Logger().Error("failed to find portal url with Lemon Squeezy api", err, appcontext.Fields{})
		return nil, err
	}
	if url == nil {
		ctx.Logger().ErrorText("portal url not found, respond")
		return nil, nil
	}

	ctx.Logger().Text("persist portal url in caching")
	if err = s.cachingRepository.SetLemonsqueezyCustomerPortalURL(ctx, customerID, *url); err != nil {
		ctx.Logger().Error("failed to persist portal url in caching", err, appcontext.Fields{})
	}

	return url, nil
}
