package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	Email     string             `bson:"email"`
	Feedback  string             `bson:"feedback"`
	Ip        string             `bson:"ip"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (f Feedback) ToDomain() domain.Feedback {
	return domain.Feedback{
		ID:        f.ID.Hex(),
		UserID:    f.UserID.Hex(),
		Email:     f.Email,
		Feedback:  f.Feedback,
		Ip:        f.Ip,
		CreatedAt: f.CreatedAt,
	}
}

func (Feedback) FromDomain(feedback domain.Feedback) (*Feedback, error) {
	id, err := database.ObjectIDFromString(feedback.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidFeedback
	}

	uid, err := database.ObjectIDFromString(feedback.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	return &Feedback{
		ID:        id,
		UserID:    uid,
		Email:     feedback.Email,
		Feedback:  feedback.Feedback,
		Ip:        feedback.Ip,
		CreatedAt: feedback.CreatedAt,
	}, nil
}
