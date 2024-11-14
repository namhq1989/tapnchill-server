package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type FeedbackRepository interface {
	Create(ctx *appcontext.AppContext, feedback Feedback) error
}

type Feedback struct {
	ID        string
	UserID    string
	Email     string
	Feedback  string
	Ip        string
	CreatedAt time.Time
}

func NewFeedback(userID, email, feedback, ip string) (*Feedback, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if len(feedback) == 0 {
		return nil, apperrors.Common.InvalidFeedback
	}

	return &Feedback{
		ID:        database.NewStringID(),
		UserID:    userID,
		Email:     email,
		Feedback:  feedback,
		Ip:        ip,
		CreatedAt: manipulation.NowUTC(),
	}, nil
}
