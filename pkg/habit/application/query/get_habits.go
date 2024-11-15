package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type GetHabitsHandler struct {
	habitRepository domain.HabitRepository
}

func NewGetHabitsHandler(habitRepository domain.HabitRepository) GetHabitsHandler {
	return GetHabitsHandler{
		habitRepository: habitRepository,
	}
}

func (h GetHabitsHandler) GetHabits(_ *appcontext.AppContext, _ string, _ dto.GetHabitsRequest) (*dto.GetHabitsResponse, error) {
	return nil, nil
}
