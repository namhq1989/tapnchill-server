package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type QRCodeRepository interface {
	Create(ctx *appcontext.AppContext, qrCode QRCode) error
	Update(ctx *appcontext.AppContext, qrCode QRCode) error
	Delete(ctx *appcontext.AppContext, qrCodeID string) error
	CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindByID(ctx *appcontext.AppContext, qrCodeID string) (*QRCode, error)
	FindByFilter(ctx *appcontext.AppContext, filter QRCodeFilter) ([]QRCode, error)
}

type QRCode struct {
	ID        string
	UserID    string
	Name      string
	Type      string
	Content   string
	Settings  QRCodeSettings
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewQRCode(userID string, name, codeType, content string, settings QRCodeSettings, data string) (*QRCode, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if codeType == "" {
		return nil, apperrors.Common.InvalidType
	}

	if content == "" {
		return nil, apperrors.Common.InvalidContent
	}

	code := &QRCode{
		ID:        database.NewStringID(),
		UserID:    userID,
		Type:      codeType,
		Content:   content,
		Settings:  settings,
		Data:      data,
		CreatedAt: manipulation.NowUTC(),
		UpdatedAt: manipulation.NowUTC(),
	}

	if err := code.SetName(name); err != nil {
		return nil, err
	}

	return code, nil
}

func (q *QRCode) SetName(name string) error {
	if len(name) < 2 {
		return apperrors.Common.InvalidName
	}

	q.Name = name
	q.SetUpdatedAt()
	return nil
}

func (q *QRCode) SetUpdatedAt() {
	q.UpdatedAt = manipulation.NowUTC()
}
