package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Habit struct {
	ID                    primitive.ObjectID `bson:"_id"`
	UserID                primitive.ObjectID `bson:"userId"`
	Name                  string             `bson:"name"`
	Goal                  string             `bson:"goal"`
	DaysOfWeek            []int              `bson:"daysOfWeek"`
	Icon                  string             `bson:"icon"`
	SortOrder             int                `bson:"sortOrder"`
	Status                string             `bson:"status"`
	StatsLongestStreak    int                `bson:"statsLongestStreak"`
	StatsCurrentStreak    int                `bson:"statsCurrentStreak"`
	StatsTotalCompletions int                `bson:"statsTotalCompletions"`
	CreatedAt             time.Time          `bson:"createdAt"`
	LastCompletedAt       *time.Time         `bson:"lastCompletedAt"`
	LastActivatedAt       time.Time          `bson:"lastActivatedAt"`
}

func (h Habit) ToDomain() domain.Habit {
	return domain.Habit{
		ID:                    h.ID.Hex(),
		UserID:                h.UserID.Hex(),
		Name:                  h.Name,
		Goal:                  h.Goal,
		DaysOfWeek:            h.DaysOfWeek,
		Icon:                  h.Icon,
		SortOrder:             h.SortOrder,
		Status:                domain.HabitStatus(h.Status),
		StatsLongestStreak:    h.StatsLongestStreak,
		StatsCurrentStreak:    h.StatsCurrentStreak,
		StatsTotalCompletions: h.StatsTotalCompletions,
		CreatedAt:             h.CreatedAt,
		LastCompletedAt:       h.LastCompletedAt,
		LastActivatedAt:       h.LastActivatedAt,
	}
}

func (Habit) FromDomain(habit domain.Habit) (*Habit, error) {
	id, err := database.ObjectIDFromString(habit.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(habit.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &Habit{
		ID:                    id,
		UserID:                uid,
		Name:                  habit.Name,
		Goal:                  habit.Goal,
		DaysOfWeek:            habit.DaysOfWeek,
		Icon:                  habit.Icon,
		SortOrder:             habit.SortOrder,
		Status:                string(habit.Status),
		StatsLongestStreak:    habit.StatsLongestStreak,
		StatsCurrentStreak:    habit.StatsCurrentStreak,
		StatsTotalCompletions: habit.StatsTotalCompletions,
		CreatedAt:             habit.CreatedAt,
		LastCompletedAt:       habit.LastCompletedAt,
		LastActivatedAt:       habit.LastActivatedAt,
	}, nil
}
