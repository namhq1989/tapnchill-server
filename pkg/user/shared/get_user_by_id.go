package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
)

func (s Service) GetUserByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	ctx.Logger().Info("get user by id", appcontext.Fields{"userID": userID})

	ctx.Logger().Text("find in caching")
	user, err := s.cachingRepository.GetUserByID(ctx, userID)
	if user != nil {
		ctx.Logger().Text("found in caching, respond")
		return user, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("find in db")
	user, err = s.userRepository.FindByID(ctx, userID)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetUserByID(ctx, userID, *user); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
	}

	return user, nil
}
