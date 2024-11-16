package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type GetHabitsHandler struct {
	service domain.Service
}

func NewGetHabitsHandler(service domain.Service) GetHabitsHandler {
	return GetHabitsHandler{
		service: service,
	}
}

func (h GetHabitsHandler) GetHabits(ctx *appcontext.AppContext, performerID string, _ dto.GetHabitsRequest) (*dto.GetHabitsResponse, error) {
	ctx.Logger().Info("new get habits request", appcontext.Fields{"performerID": performerID})

	habits, err := h.service.GetUserHabits(ctx, performerID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.Habit, 0)
	for _, habit := range habits {
		result = append(result, dto.Habit{}.FromDomain(habit))
	}

	ctx.Logger().Text("done get habits request")
	return &dto.GetHabitsResponse{
		Habits: result,
	}, nil
}
