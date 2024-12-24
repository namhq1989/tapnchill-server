package domain

import (
	"sort"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
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

	sort.Ints(daysOfWeek)
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

func (h *Habit) IsPreviousScheduledDayOf(currentDate time.Time) bool {
	currentWeekday := int(currentDate.Weekday())

	isValidDay := false
	for _, day := range h.DaysOfWeek {
		if day == currentWeekday {
			isValidDay = true
			break
		}
	}
	if !isValidDay {
		return false
	}

	if len(h.DaysOfWeek) == 1 {
		if currentWeekday == h.DaysOfWeek[0] {
			previousScheduledDate := currentDate.AddDate(0, 0, -7)
			return h.LastCompletedAt.Equal(previousScheduledDate)
		}
	}

	currentIndex := -1
	for i, day := range h.DaysOfWeek {
		if day == currentWeekday {
			currentIndex = i
			break
		}
	}
	previousIndex := (currentIndex - 1 + len(h.DaysOfWeek)) % len(h.DaysOfWeek)
	previousScheduledDay := h.DaysOfWeek[previousIndex]

	offset := (currentWeekday - previousScheduledDay + 7) % 7
	lastWeekStart := currentDate.AddDate(0, 0, -int(currentDate.Weekday()))
	lastDoneWeekStart := h.LastCompletedAt.AddDate(0, 0, -int(h.LastCompletedAt.Weekday()))
	if lastDoneWeekStart.Before(lastWeekStart) {
		offset += 1
	}

	previousScheduledDate := currentDate.AddDate(0, 0, -offset)
	isEqual := h.LastCompletedAt.Year() == previousScheduledDate.Year() &&
		h.LastCompletedAt.Month() == previousScheduledDate.Month() &&
		h.LastCompletedAt.Day() == previousScheduledDate.Day()

	return isEqual
}

func (h *Habit) OnCompleted(date time.Time) {
	tz := manipulation.GetTimezoneIdentifier(date)
	if h.LastCompletedAt != nil && manipulation.IsSameDay(*h.LastCompletedAt, date, tz) {
		return
	}

	// update total completions
	h.StatsTotalCompletions++

	// if no previous completion, initialize streaks
	if h.LastCompletedAt == nil {
		h.LastCompletedAt = &date
		h.StatsCurrentStreak = 1
		h.StatsLongestStreak = 1
		return
	}

	var (
		today           = manipulation.Now(tz)
		lastCompletedAt = h.LastCompletedAt.In(today.Location())
	)
	if !manipulation.IsSameDay(date, today, tz) {
		// if not today
		if lastCompletedAt.Before(date) {
			h.LastCompletedAt = &date
			h.StatsCurrentStreak = 1
			if h.StatsLongestStreak == 0 {
				h.StatsLongestStreak = 1
			}
		}
		return
	}

	if isPreviousScheduledDay := h.IsPreviousScheduledDayOf(date); !isPreviousScheduledDay {
		h.StatsCurrentStreak = 1
		if h.StatsLongestStreak == 0 {
			h.StatsLongestStreak = 1
		}
	} else {
		h.StatsCurrentStreak++
		if h.StatsCurrentStreak > h.StatsLongestStreak {
			h.StatsLongestStreak = h.StatsCurrentStreak
		}
	}
	h.LastCompletedAt = &date
}

func (h *Habit) IsActive() bool {
	return h.Status == HabitStatusActive
}

func (h *Habit) IsInactive() bool {
	return h.Status == HabitStatusInactive
}

type HabitFilter struct {
	UserID string
}

func NewHabitFilter(userID string) (*HabitFilter, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	return &HabitFilter{
		UserID: userID,
	}, nil
}
