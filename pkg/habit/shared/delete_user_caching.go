package shared

import (
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

func (s Service) DeleteUserCaching(ctx *appcontext.AppContext, userID string, date time.Time) error {
	ctx.Logger().Text("get start of day")
	startOfDay := manipulation.StartOfDay(date)

	ctx.Logger().Text("delete habits caching")
	if err := s.cachingRepository.DeleteUserHabits(ctx, userID); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	dateDDMM := fmt.Sprintf("%02d%02d", startOfDay.Day(), int(startOfDay.Month()))
	ctx.Logger().Info("delete stats caching", appcontext.Fields{"dateDDMM": dateDDMM})
	if err := s.cachingRepository.DeleteUserStats(ctx, userID, dateDDMM); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	return nil
}
