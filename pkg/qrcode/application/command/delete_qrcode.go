package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

type DeleteQRCodeHandler struct {
	qrCodeRepository domain.QRCodeRepository
}

func NewDeleteQRCodeHandler(qrCodeRepository domain.QRCodeRepository) DeleteQRCodeHandler {
	return DeleteQRCodeHandler{
		qrCodeRepository: qrCodeRepository,
	}
}

func (h DeleteQRCodeHandler) DeleteQRCode(ctx *appcontext.AppContext, performerID, qrCodeID string, _ dto.DeleteQRCodeRequest) (*dto.DeleteQRCodeResponse, error) {
	ctx.Logger().Info("new delete qr code request", appcontext.Fields{
		"performerID": performerID, "qrCodeID": qrCodeID,
	})

	qrCode, err := h.qrCodeRepository.FindByID(ctx, qrCodeID)
	if err != nil {
		ctx.Logger().Error("failed to find QR code in db", err, appcontext.Fields{})
		return nil, err
	}
	if qrCode == nil {
		ctx.Logger().ErrorText("QR code not found, respond")
		return nil, apperrors.Common.NotFound
	}
	if qrCode.UserID != performerID {
		ctx.Logger().ErrorText("QR code author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("delete QR code in db")
	if err = h.qrCodeRepository.Delete(ctx, qrCode.ID); err != nil {
		ctx.Logger().Error("failed to delete QR code in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done delete QR code request")
	return &dto.DeleteQRCodeResponse{}, nil
}
