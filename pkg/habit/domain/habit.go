package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HabitRepository interface {
	Create(ctx *appcontext.AppContext, habit Habit) error
	Update(ctx *appcontext.AppContext, habit Habit) error
	Delete(ctx *appcontext.AppContext, habitID string) error
	CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindByID(ctx *appcontext.AppContext, habitID string) (*Habit, error)
	FindByFilter(ctx *appcontext.AppContext, filter HabitFilter) ([]Habit, error)
}

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
	LastCompletedAt       *time.Time
	LastActivatedAt       time.Time
}

func NewHabit(userID, name, goal string, daysOfWeek []int, icon string, sortOrder int) (*Habit, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	habit := Habit{
		ID:              database.NewStringID(),
		UserID:          userID,
		SortOrder:       sortOrder,
		Status:          HabitStatusActive,
		CreatedAt:       manipulation.NowUTC(),
		LastActivatedAt: manipulation.NowUTC(),
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

func (h *Habit) SetStatus(status HabitStatus) {
	h.Status = status

	if h.IsActive() {
		h.LastActivatedAt = manipulation.NowUTC()
	}
}

func (h *Habit) SetSortOrder(order int) {
	h.SortOrder = order
}

func (h *Habit) OnCompleted(date time.Time) {
	var (
		tz = manipulation.GetTimezoneIdentifier(date)
	)

	var previousDate time.Time
	if h.LastCompletedAt != nil {
		previousDate = *h.LastCompletedAt
	} else {
		previousDate = time.Time{}
	}

	expectedPreviousDate := manipulation.PreviousDay(date, tz)
	if manipulation.IsSameDay(previousDate, expectedPreviousDate, tz) {
		h.StatsCurrentStreak++
	} else {
		h.StatsCurrentStreak = 1
	}

	if h.LastCompletedAt == nil || date.After(*h.LastCompletedAt) {
		h.LastCompletedAt = &date
	}
	h.StatsTotalCompletions++

	if h.StatsCurrentStreak > h.StatsLongestStreak {
		h.StatsLongestStreak = h.StatsCurrentStreak
	}
}

func (h *Habit) IsActive() bool {
	return h.Status == HabitStatusActive
}

func (h *Habit) IsInactive() bool {
	return h.Status == HabitStatusInactive
}

type HabitFilter struct {
	UserID primitive.ObjectID
}

func NewHabitFilter(userID string) (*HabitFilter, error) {
	uid, err := database.ObjectIDFromString(userID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitFilter{
		UserID: uid,
	}, nil
}
