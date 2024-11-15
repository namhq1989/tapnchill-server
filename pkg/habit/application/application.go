package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
)

type (
	Commands interface {
		CreateHabit(ctx *appcontext.AppContext, performerID string, req dto.CreateHabitRequest) (*dto.CreateHabitResponse, error)
		UpdateHabit(ctx *appcontext.AppContext, performerID, habitID string, req dto.UpdateHabitRequest) (*dto.UpdateHabitResponse, error)
		ChangeHabitStatus(ctx *appcontext.AppContext, performerID, habitID string, _ dto.ChangeHabitStatusRequest) (*dto.ChangeHabitStatusResponse, error)
		CompleteHabit(ctx *appcontext.AppContext, performerID, habitID string, _ dto.CompleteHabitRequest) (*dto.CompleteHabitResponse, error)
	}
	Queries interface {
		GetHabits(_ *appcontext.AppContext, performerID string, _ dto.GetHabitsRequest) (*dto.GetHabitsResponse, error)
		GetTodayHabits(_ *appcontext.AppContext, performerID string, _ dto.GetTodayHabitsRequest) (*dto.GetTodayHabitsResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.CreateHabitHandler
		command.UpdateHabitHandler
		command.ChangeHabitStatusHandler
		command.CompleteHabitHandler
	}
	queryHandlers struct {
		query.GetHabitsHandler
		query.GetTodayHabitsHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	habitRepository domain.HabitRepository,
	habitCompletionRepository domain.HabitCompletionRepository,
	habitDailyStatsRepository domain.HabitDailyStatsRepository,
	service domain.Service,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateHabitHandler:       command.NewCreateHabitHandler(habitRepository),
			UpdateHabitHandler:       command.NewUpdateHabitHandler(habitRepository, service),
			ChangeHabitStatusHandler: command.NewChangeHabitStatusHandler(habitRepository, service),
			CompleteHabitHandler:     command.NewCompleteHabitHandler(habitRepository, habitCompletionRepository, habitDailyStatsRepository, service),
		},
		queryHandlers: queryHandlers{
			GetHabitsHandler:      query.NewGetHabitsHandler(habitRepository),
			GetTodayHabitsHandler: query.NewGetTodayHabitsHandler(habitRepository),
		},
	}
}
