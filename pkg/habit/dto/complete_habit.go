package dto

import "github.com/namhq1989/tapnchill-server/internal/utils/httprespond"

type CompleteHabitRequest struct {
	Date string `json:"date" validate:"required" message:"invalid_date"`
}

type CompleteHabitResponse struct {
	LastCompletedAt       *httprespond.TimeResponse `json:"lastCompletedAt"`
	StatsTotalCompletions int                       `json:"statsTotalCompletions"`
	StatsLongestStreak    int                       `json:"statsLongestStreak"`
	StatsCurrentStreak    int                       `json:"statsCurrentStreak"`
}
