package shared

import "github.com/namhq1989/tapnchill-server/pkg/habit/domain"

type Service struct {
	habitRepository domain.HabitRepository
}

func NewService(habitRepository domain.HabitRepository) Service {
	return Service{
		habitRepository: habitRepository,
	}
}
