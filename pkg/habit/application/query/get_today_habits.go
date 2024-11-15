package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type GetTodayHabitsHandler struct {
	habitRepository domain.HabitRepository
}

func NewGetTodayHabitsHandler(habitRepository domain.HabitRepository) GetTodayHabitsHandler {
	return GetTodayHabitsHandler{
		habitRepository: habitRepository,
	}
}

func (h GetTodayHabitsHandler) GetTodayHabits(_ *appcontext.AppContext, _ string, _ dto.GetTodayHabitsRequest) (*dto.GetTodayHabitsResponse, error) {
	return nil, nil
}
