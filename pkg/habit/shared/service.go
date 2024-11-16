package shared

import "github.com/namhq1989/tapnchill-server/pkg/habit/domain"

type Service struct {
	habitRepository           domain.HabitRepository
	habitDailyStatsRepository domain.HabitDailyStatsRepository
	cachingRepository         domain.CachingRepository
}

func NewService(habitRepository domain.HabitRepository, habitDailyStatsRepository domain.HabitDailyStatsRepository, cachingRepository domain.CachingRepository) Service {
	return Service{
		habitRepository:           habitRepository,
		habitDailyStatsRepository: habitDailyStatsRepository,
		cachingRepository:         cachingRepository,
	}
}
