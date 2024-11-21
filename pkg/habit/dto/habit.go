package dto

import (
	"time"

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
	LastActivatedAt       *httprespond.TimeResponse `json:"lastActivatedAt"`
}

func (Habit) FromDomain(habit domain.Habit) Habit {
	h := Habit{
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
		LastActivatedAt:       httprespond.NewTimeResponse(habit.LastActivatedAt),
	}

	if habit.LastCompletedAt != nil {
		h.LastCompletedAt = httprespond.NewTimeResponse(*habit.LastCompletedAt)
	} else {
		h.LastCompletedAt = httprespond.NewTimeResponse(time.Time{})
	}

	return h
}
