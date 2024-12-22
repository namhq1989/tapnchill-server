package dto

import (
	"time"

	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
)

type QRCode struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Content   string         `json:"content"`
	Settings  QRCodeSettings `json:"settings"`
	Data      string         `json:"data"`
	CreatedAt time.Time      `json:"createdAt"`
}

func (QRCode) FromDomain(code domain.QRCode) QRCode {
	return QRCode{
		ID:        code.ID,
		Name:      code.Name,
		Type:      code.Type,
		Content:   code.Content,
		Settings:  QRCodeSettings{}.FromDomain(code.Settings),
		Data:      code.Data,
		CreatedAt: code.CreatedAt,
	}
}
