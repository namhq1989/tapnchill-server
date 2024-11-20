package dto

import (
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
)

type Habit struct {
	ID                    string                    `json:"id"`
	Name                  string                    `json:"name"`
	Goal                  string                    `json:"goal"`
	DaysOfWeek            []int                     `json:"daysOfWeek"`
	Icon                  string                    `json:"icon"`
	SortOrder             int                       `json:"sortOrder"`
	Status                string                    `json:"status"`
	StatsLongestStreak    int                       `json:"statsLongestStreak"`
	StatsCurrentStreak    int                       `json:"statsCurrentStreak"`
	StatsTotalCompletions int                       `json:"statsTotalCompletions"`
	CreatedAt             *httprespond.TimeResponse `json:"createdAt"`
	LastCompletedAt       *httprespond.TimeResponse `json:"lastCompletedAt"`
	LastActivatedAt       *httprespond.TimeResponse `json:"LastActivatedAt"`
}

func (Habit) FromDomain(habit domain.Habit) Habit {
	return Habit{
		ID:                    habit.ID,
		Name:                  habit.Name,
		Goal:                  habit.Goal,
		DaysOfWeek:            habit.DaysOfWeek,
		Icon:                  habit.Icon,
		SortOrder:             habit.SortOrder,
		Status:                string(habit.Status),
		StatsLongestStreak:    habit.StatsLongestStreak,
		StatsCurrentStreak:    habit.StatsCurrentStreak,
		StatsTotalCompletions: habit.StatsTotalCompletions,
		CreatedAt:             httprespond.NewTimeResponse(habit.CreatedAt),
		LastCompletedAt:       httprespond.NewTimeResponse(habit.LastCompletedAt),
		LastActivatedAt:       httprespond.NewTimeResponse(habit.LastActivatedAt),
	}
}
