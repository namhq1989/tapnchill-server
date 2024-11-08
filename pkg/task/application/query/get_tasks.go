package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type GetTasksHandler struct {
	taskRepository domain.TaskRepository
}

func NewGetTasksHandler(taskRepository domain.TaskRepository) GetTasksHandler {
	return GetTasksHandler{
		taskRepository: taskRepository,
	}
}

func (h GetTasksHandler) GetTasks(ctx *appcontext.AppContext, performerID string, req dto.GetTasksRequest) (*dto.GetTasksResponse, error) {
	ctx.Logger().Info("new get tasks request", appcontext.Fields{
		"performerID": performerID, "goalId": req.GoalID,
		"keyword": req.Keyword, "status": req.Status, "pageToken": req.PageToken,
	})

	ctx.Logger().Text("create filter")
	filter, err := domain.NewTaskFilter(performerID, req.GoalID, req.Status, req.Keyword, req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to create filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find tasks in db")
	tasks, err := h.taskRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find tasks in db", err, appcontext.Fields{})
		return nil, err
	}

	totalTasks := len(tasks)
	if totalTasks == 0 {
		ctx.Logger().Text("no tasks found, respond")
		return &dto.GetTasksResponse{
			Tasks:         make([]dto.Task, 0),
			NextPageToken: "",
		}, nil
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.Task, 0)
	for _, task := range tasks {
		result = append(result, dto.Task{}.FromDomain(task))
	}

	nextPageToken := ""
	if totalTasks == int(filter.Limit) {
		nextPageToken = pagetoken.NewWithTimestamp(tasks[totalTasks-1].CreatedAt)
	}

	ctx.Logger().Text("done get tasks request")
	return &dto.GetTasksResponse{
		Tasks:         result,
		NextPageToken: nextPageToken,
	}, nil
}
