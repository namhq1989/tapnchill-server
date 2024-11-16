package infrastructure

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/caching"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

type CachingRepository struct {
	caching caching.Operations

	domain             string
	userHabitsDuration time.Duration
	userStatsDuration  time.Duration
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching:            caching,
		domain:             "habit",
		userHabitsDuration: 24 * time.Hour,
		userStatsDuration:  24 * time.Hour,
	}
}

// USER HABITS

func (r CachingRepository) GetUserHabits(ctx *appcontext.AppContext, userID string) ([]domain.Habit, error) {
	key := r.generateUserHabitsKey(userID)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result []domain.Habit
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r CachingRepository) SetUserHabits(ctx *appcontext.AppContext, userID string, habits []domain.Habit) error {
	key := r.generateUserHabitsKey(userID)
	r.caching.SetTTL(ctx, key, habits, r.userHabitsDuration)
	return nil
}

func (r CachingRepository) DeleteUserHabits(ctx *appcontext.AppContext, userID string) error {
	key := r.generateUserHabitsKey(userID)
	_, _ = r.caching.Del(ctx, key)
	return nil
}

func (r CachingRepository) generateUserHabitsKey(userID string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("user:%s", userID))
}

// USER STATS

func (r CachingRepository) GetUserStats(ctx *appcontext.AppContext, userID, date string) ([]domain.HabitDailyStats, error) {
	key := r.generateUserStatsKey(userID, date)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result []domain.HabitDailyStats
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return result, nil

}

func (r CachingRepository) SetUserStats(ctx *appcontext.AppContext, userID, date string, stats []domain.HabitDailyStats) error {
	key := r.generateUserStatsKey(userID, date)
	r.caching.SetTTL(ctx, key, stats, r.userStatsDuration)
	return nil
}

func (r CachingRepository) DeleteUserStats(ctx *appcontext.AppContext, userID, date string) error {
	key := r.generateUserStatsKey(userID, date)
	_, _ = r.caching.Del(ctx, key)
	return nil
}

func (r CachingRepository) generateUserStatsKey(userID, date string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("user:%s:%s:stats", userID, date))
}
