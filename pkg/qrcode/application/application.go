package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

type (
	Commands interface {
		CreateQRCode(ctx *appcontext.AppContext, performerID string, req dto.CreateQRCodeRequest) (*dto.CreateQRCodeResponse, error)
		UpdateQRCode(ctx *appcontext.AppContext, performerID, qrCodeID string, req dto.UpdateQRCodeRequest) (*dto.UpdateQRCodeResponse, error)
		DeleteQRCode(ctx *appcontext.AppContext, performerID, qrCodeID string, _ dto.DeleteQRCodeRequest) (*dto.DeleteQRCodeResponse, error)
	}
	Queries interface {
		GetQRCodes(ctx *appcontext.AppContext, performerID string, req dto.GetQRCodesRequest) (*dto.GetQRCodesResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.CreateQRCodeHandler
		command.UpdateQRCodeHandler
		command.DeleteQRCodeHandler
	}
	queryHandlers struct {
		query.GetQRCodesHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	qrCodeRepository domain.QRCodeRepository,
	userHub domain.UserHub,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateQRCodeHandler: command.NewCreateQRCodeHandler(qrCodeRepository, userHub),
			UpdateQRCodeHandler: command.NewUpdateQRCodeHandler(qrCodeRepository),
			DeleteQRCodeHandler: command.NewDeleteQRCodeHandler(qrCodeRepository),
		},
		queryHandlers: queryHandlers{
			GetQRCodesHandler: query.NewGetQRCodesHandler(qrCodeRepository),
		},
	}
}
