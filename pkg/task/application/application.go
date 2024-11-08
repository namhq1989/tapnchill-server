package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
)

type (
	Commands interface {
		CreateGoal(ctx *appcontext.AppContext, performerID string, req dto.CreateGoalRequest) (*dto.CreateGoalResponse, error)
		UpdateGoal(ctx *appcontext.AppContext, performerID, goalID string, req dto.UpdateGoalRequest) (*dto.UpdateGoalResponse, error)
		DeleteGoal(ctx *appcontext.AppContext, performerID, goalID string, _ dto.DeleteGoalRequest) (*dto.DeleteGoalResponse, error)

		CreateTask(ctx *appcontext.AppContext, performerID string, req dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
		UpdateTask(ctx *appcontext.AppContext, performerID, taskID string, req dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error)
		ChangeTaskStatus(ctx *appcontext.AppContext, performerID, taskID string, req dto.ChangeTaskStatusRequest) (*dto.ChangeTaskStatusResponse, error)
		DeleteTask(ctx *appcontext.AppContext, performerID, taskID string, _ dto.DeleteTaskRequest) (*dto.DeleteTaskResponse, error)
	}
	Queries interface {
		GetGoals(ctx *appcontext.AppContext, performerID string, req dto.GetGoalsRequest) (*dto.GetGoalsResponse, error)
		GetTasks(ctx *appcontext.AppContext, performerID string, req dto.GetTasksRequest) (*dto.GetTasksResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.CreateGoalHandler
		command.UpdateGoalHandler
		command.DeleteGoalHandler

		command.CreateTaskHandler
		command.UpdateTaskHandler
		command.ChangeTaskStatusHandler
		command.DeleteTaskHandler
	}
	queryHandlers struct {
		query.GetGoalsHandler
		query.GetTasksHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	taskRepository domain.TaskRepository,
	goalRepository domain.GoalRepository,
	service domain.Service,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateGoalHandler: command.NewCreateGoalHandler(goalRepository),
			UpdateGoalHandler: command.NewUpdateGoalHandler(goalRepository, service),
			DeleteGoalHandler: command.NewDeleteGoalHandler(goalRepository, service),

			CreateTaskHandler: command.NewCreateTaskHandler(taskRepository, goalRepository, service),
			UpdateTaskHandler: command.NewUpdateTaskHandler(taskRepository, service),
			ChangeTaskStatusHandler: command.NewChangeTaskStatusHandler(
				taskRepository,
				goalRepository,
				service,
			),
			DeleteTaskHandler: command.NewDeleteTaskHandler(taskRepository, goalRepository, service),
		},
		queryHandlers: queryHandlers{
			GetGoalsHandler: query.NewGetGoalsHandler(goalRepository),
			GetTasksHandler: query.NewGetTasksHandler(taskRepository),
		},
	}
}
