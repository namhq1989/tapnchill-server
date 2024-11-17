package shared

import (
	"fmt"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

func (s Service) DeleteUserCaching(ctx *appcontext.AppContext, userID, date string) error {
	ctx.Logger().Text("get start of day with client date")
	startOfDay, err := manipulation.GetStartOfDayWithClientDate(date)
	if err != nil {
		ctx.Logger().Error("failed to get start of day with client date", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("delete habits caching")
	if err = s.cachingRepository.DeleteUserHabits(ctx, userID); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	dateDDMM := fmt.Sprintf("%02d%02d", startOfDay.Day(), int(startOfDay.Month()))
	ctx.Logger().Info("delete stats caching", appcontext.Fields{"dateDDMM": dateDDMM})
	if err = s.cachingRepository.DeleteUserStats(ctx, userID, dateDDMM); err != nil {
		ctx.Logger().Error("failed to delete in caching", err, appcontext.Fields{})
	}

	return nil
}
