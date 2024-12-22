package dbmodel

import (
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
)

type QRCodeSettings struct {
	Color    string `bson:"color"`
	HasLogo  bool   `bson:"hasLogo"`
	LogoData string `bson:"logoData"`
	LogoName string `bson:"logoName"`
	Style    string `bson:"style"`
}

func (s QRCodeSettings) ToDomain() domain.QRCodeSettings {
	return domain.QRCodeSettings{
		Color:    s.Color,
		HasLogo:  s.HasLogo,
		LogoData: s.LogoData,
		LogoName: s.LogoName,
		Style:    s.Style,
	}
}

func (QRCodeSettings) FromDomain(settings domain.QRCodeSettings) (*QRCodeSettings, error) {
	return &QRCodeSettings{
		Color:    settings.Color,
		HasLogo:  settings.HasLogo,
		LogoData: settings.LogoData,
		LogoName: settings.LogoName,
		Style:    settings.Style,
	}, nil
}
