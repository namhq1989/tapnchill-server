package domain

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type Habit struct {
	ID                    string
	UserID                string
	Name                  string
	Goal                  string
	DaysOfWeek            []int
	Icon                  string
	SortOrder             int
	Status                HabitStatus
	StatsLongestStreak    int
	StatsCurrentStreak    int
	StatsTotalCompletions int
	CreatedAt             time.Time
	LastCompletedAt       time.Time
}

func NewHabit(userID, name, goal string, daysOfWeek []int, icon string, sortOrder int) (*Habit, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	habit := Habit{
		ID:        database.NewStringID(),
		UserID:    userID,
		SortOrder: sortOrder,
		Status:    HabitStatusActive,
		CreatedAt: manipulation.NowUTC(),
	}

	if err := habit.SetName(name); err != nil {
		return nil, err
	}
	if err := habit.SetGoal(goal); err != nil {
		return nil, err
	}
	if err := habit.SetDaysOfWeek(daysOfWeek); err != nil {
		return nil, err
	}
	if err := habit.SetIcon(icon); err != nil {
		return nil, err
	}

	return &habit, nil
}

func (h *Habit) SetName(name string) error {
	if len(name) < 2 || len(name) > 30 {
		return apperrors.Common.InvalidName
	}

	h.Name = name
	return nil
}

func (h *Habit) SetGoal(goal string) error {
	if len(goal) < 2 || len(goal) > 50 {
		return apperrors.Common.InvalidGoal
	}

	h.Goal = goal
	return nil
}

func (h *Habit) SetDaysOfWeek(daysOfWeek []int) error {
	if len(daysOfWeek) < 1 || len(daysOfWeek) > 7 {
		return apperrors.Common.InvalidDaysOfWeek
	}

	h.DaysOfWeek = daysOfWeek
	return nil
}

func (h *Habit) SetIcon(icon string) error {
	if icon == "" {
		return apperrors.Common.InvalidIcon
	}

	h.Icon = icon
	return nil
}
