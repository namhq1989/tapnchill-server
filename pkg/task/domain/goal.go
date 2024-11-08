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
	TotalTask          int
	TotalCompletedTask int
}

func NewGoal(userID, name, description string) (*Goal, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if len(name) < 2 {
		return nil, apperrors.Common.InvalidName
	}

	return &Goal{
		ID:           database.NewStringID(),
		UserID:       userID,
		Name:         name,
		Description:  description,
		SearchString: manipulation.NormalizeText(fmt.Sprintf("%s %s", name, description)),
		Stats: GoalStats{
			TotalTask:          0,
			TotalCompletedTask: 0,
		},
		IsCompleted: false,
		CreatedAt:   manipulation.NowUTC(),
	}, nil
}

func (g *Goal) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	g.Name = name
	g.SearchString = manipulation.NormalizeText(fmt.Sprintf("%s %s", name, g.Description))
	g.SetUpdatedAt()
	return nil
}

func (g *Goal) SetDescription(description string) {
	g.Description = description
	g.SearchString = manipulation.NormalizeText(fmt.Sprintf("%s %s", g.Name, description))
	g.SetUpdatedAt()
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

func (g *Goal) AdjustTotalCompletedTask(value int) {
	g.Stats.TotalCompletedTask += value
	if g.Stats.TotalCompletedTask < 0 {
		g.Stats.TotalCompletedTask = 0
	}
	g.SetUpdatedAt()
}

func (g *Goal) SetUpdatedAt() {
	g.UpdatedAt = manipulation.NowUTC()
}
