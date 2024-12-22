package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

type UpdateQRCodeHandler struct {
	qrCodeRepository domain.QRCodeRepository
}

func NewUpdateQRCodeHandler(
	qrCodeRepository domain.QRCodeRepository,
) UpdateQRCodeHandler {
	return UpdateQRCodeHandler{
		qrCodeRepository: qrCodeRepository,
	}
}

func (h UpdateQRCodeHandler) UpdateQRCode(ctx *appcontext.AppContext, performerID, qrCodeID string, req dto.UpdateQRCodeRequest) (*dto.UpdateQRCodeResponse, error) {
	ctx.Logger().Info("new update QR code request", appcontext.Fields{
		"performerID": performerID, "qrCodeID": qrCodeID, "name": req.Name,
	})

	ctx.Logger().Text("find QR code in db")
	qrCode, err := h.qrCodeRepository.FindByID(ctx, qrCodeID)
	if err != nil {
		ctx.Logger().Error("failed to find Qr code in db", err, appcontext.Fields{})
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

	ctx.Logger().Text("update QR code data")
	if err = qrCode.SetName(req.Name); err != nil {
		ctx.Logger().Error("failed to update QR code name", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("update QR code in db")
	if err = h.qrCodeRepository.Update(ctx, *qrCode); err != nil {
		ctx.Logger().Error("failed to update QR code in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update QR code request")
	return &dto.UpdateQRCodeResponse{
		ID: qrCode.ID,
	}, nil
}
