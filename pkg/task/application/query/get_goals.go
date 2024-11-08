package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type GetGoalsHandler struct {
	goalRepository domain.GoalRepository
}

func NewGetGoalsHandler(goalRepository domain.GoalRepository) GetGoalsHandler {
	return GetGoalsHandler{
		goalRepository: goalRepository,
	}
}

func (h GetGoalsHandler) GetGoals(ctx *appcontext.AppContext, performerID string, req dto.GetGoalsRequest) (*dto.GetGoalsResponse, error) {
	ctx.Logger().Info("new get goals request", appcontext.Fields{"performerID": performerID, "keyword": req.Keyword, "pageToken": req.PageToken})

	ctx.Logger().Text("create filter")
	filter, err := domain.NewGoalFilter(performerID, req.Keyword, req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to create filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find goals in db")
	goals, err := h.goalRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find goals in db", err, appcontext.Fields{})
		return nil, err
	}

	totalGoals := len(goals)
	if totalGoals == 0 {
		ctx.Logger().Text("no goals found, respond")
		return &dto.GetGoalsResponse{
			Goals:         make([]dto.Goal, 0),
			NextPageToken: "",
		}, nil
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.Goal, 0)
	for _, goal := range goals {
		result = append(result, dto.Goal{}.FromDomain(goal))
	}

	nextPageToken := ""
	if totalGoals == int(filter.Limit) {
		nextPageToken = pagetoken.NewWithTimestamp(goals[totalGoals-1].CreatedAt)
	}

	ctx.Logger().Text("done get goals request")
	return &dto.GetGoalsResponse{
		Goals:         result,
		NextPageToken: nextPageToken,
	}, nil
}
