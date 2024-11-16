package shared

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

func (s Service) GetHabitByID(ctx *appcontext.AppContext, habitID, userID string) (*domain.Habit, error) {
	ctx.Logger().Text("find habit in db")
	habit, err := s.habitRepository.FindByID(ctx, habitID)
	if err != nil {
		ctx.Logger().Error("failed to find habit in db", err, appcontext.Fields{})
		return nil, err
	}
	if habit == nil {
		ctx.Logger().ErrorText("habit not found, respond")
		return nil, apperrors.Common.NotFound
	}

	if habit.UserID != userID {
		ctx.Logger().ErrorText("habit author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	return habit, nil
}
