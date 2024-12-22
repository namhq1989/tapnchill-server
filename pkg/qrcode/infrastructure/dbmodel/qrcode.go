package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QRCode struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userId"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	Content   string             `bson:"content"`
	Settings  QRCodeSettings     `bson:"settings"`
	Data      string             `bson:"data"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (q QRCode) ToDomain() domain.QRCode {
	return domain.QRCode{
		ID:        q.ID.Hex(),
		UserID:    q.UserID.Hex(),
		Name:      q.Name,
		Type:      q.Type,
		Content:   q.Content,
		Settings:  q.Settings.ToDomain(),
		Data:      q.Data,
		CreatedAt: q.CreatedAt,
		UpdatedAt: q.UpdatedAt,
	}
}

func (QRCode) FromDomain(code domain.QRCode) (*QRCode, error) {
	id, err := database.ObjectIDFromString(code.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	uid, err := database.ObjectIDFromString(code.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	settings, _ := QRCodeSettings{}.FromDomain(code.Settings)

	return &QRCode{
		ID:        id,
		UserID:    uid,
		Name:      code.Name,
		Type:      code.Type,
		Content:   code.Content,
		Settings:  *settings,
		Data:      code.Data,
		CreatedAt: code.CreatedAt,
		UpdatedAt: code.UpdatedAt,
	}, nil
}
