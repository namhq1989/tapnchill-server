package domain

import (
	"fmt"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type GoalRepository interface {
	Create(ctx *appcontext.AppContext, goal Goal) error
	Update(ctx *appcontext.AppContext, goal Goal) error
	Delete(ctx *appcontext.AppContext, id string) error
	CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindByFilter(ctx *appcontext.AppContext, filter GoalFilter) ([]Goal, error)
	FindByID(ctx *appcontext.AppContext, goalID string) (*Goal, error)
}

type Goal struct {
	ID           string
	UserID       string
	Name         string
	Description  string
	SearchString string
	Stats        GoalStats
	IsCompleted  bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type GoalStats struct {
	TotalTask     int
	TotalDoneTask int
}

const (
	DefaultGoalName        = "Basecamp"
	DefaultGoalDescription = "The starting point for your journey, where you set your sights and prepare for each step ahead"
)

func NewGoal(userID, name, description string) (*Goal, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	g := &Goal{
		ID:     database.NewStringID(),
		UserID: userID,
		Stats: GoalStats{
			TotalTask:     0,
			TotalDoneTask: 0,
		},
		IsCompleted: false,
		CreatedAt:   manipulation.NowUTC(),
	}

	if err := g.SetName(name); err != nil {
		return nil, err
	}

	if err := g.SetDescription(description); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Goal) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	g.Name = name
	g.SearchString = manipulation.NormalizeText(fmt.Sprintf("%s", name))
	g.SetUpdatedAt()
	return nil
}

func (g *Goal) SetDescription(description string) error {
	if len(description) > 1000 {
		return apperrors.Common.InvalidDescription
	}

	g.Description = description
	g.SetUpdatedAt()

	return nil
}

func (g *Goal) SetIsCompleted(isCompleted bool) {
	g.IsCompleted = isCompleted
	g.SetUpdatedAt()
}

func (g *Goal) AdjustTotalTask(value int) {
	g.Stats.TotalTask += value
	if g.Stats.TotalTask < 0 {
		g.Stats.TotalTask = 0
	}
	g.SetUpdatedAt()
}

func (g *Goal) AdjustTotalDoneTask(value int) {
	g.Stats.TotalDoneTask += value
	if g.Stats.TotalDoneTask < 0 {
		g.Stats.TotalDoneTask = 0
	}
	g.SetUpdatedAt()
}

func (g *Goal) SetUpdatedAt() {
	g.UpdatedAt = manipulation.NowUTC()
}
