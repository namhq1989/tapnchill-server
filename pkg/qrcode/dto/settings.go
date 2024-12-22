package dto

import "github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"

type QRCodeSettings struct {
	Color    string `json:"color"`
	HasLogo  bool   `json:"hasLogo"`
	LogoData string `json:"logoData"`
	LogoName string `json:"logoName"`
	Style    string `json:"style"`
}

func (QRCodeSettings) FromDomain(settings domain.QRCodeSettings) QRCodeSettings {
	return QRCodeSettings{
		Color:    settings.Color,
		HasLogo:  settings.HasLogo,
		LogoData: settings.LogoData,
		LogoName: settings.LogoName,
		Style:    settings.Style,
	}
}
