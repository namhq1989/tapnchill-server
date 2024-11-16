package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

func (s Service) GetUserHabits(ctx *appcontext.AppContext, userID string) ([]domain.Habit, error) {
	ctx.Logger().Info("get user habits", appcontext.Fields{"userID": userID})

	ctx.Logger().Text("find in caching")
	habits, err := s.cachingRepository.GetUserHabits(ctx, userID)
	if habits != nil {
		ctx.Logger().Text("found in caching, respond")
		return habits, nil
	}
	if err != nil {
		ctx.Logger().Error("failed to find in caching", err, appcontext.Fields{})
	}

	ctx.Logger().Text("create new filter")
	filter, err := domain.NewHabitFilter(userID)
	if err != nil {
		ctx.Logger().Error("failed to create new filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find in db")
	habits, err = s.habitRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist in caching")
	if err = s.cachingRepository.SetUserHabits(ctx, userID, habits); err != nil {
		ctx.Logger().Error("failed to persist in caching", err, appcontext.Fields{})
	}

	return habits, nil
}
