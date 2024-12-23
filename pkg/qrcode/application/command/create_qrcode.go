package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

type CreateQRCodeHandler struct {
	qrCodeRepository domain.QRCodeRepository
	userHub          domain.UserHub
}

func NewCreateQRCodeHandler(
	qrCodeRepository domain.QRCodeRepository,
	userHub domain.UserHub,
) CreateQRCodeHandler {
	return CreateQRCodeHandler{
		qrCodeRepository: qrCodeRepository,
		userHub:          userHub,
	}
}

func (h CreateQRCodeHandler) CreateQRCode(ctx *appcontext.AppContext, performerID string, req dto.CreateQRCodeRequest) (*dto.CreateQRCodeResponse, error) {
	ctx.Logger().Info("new create QR code request", appcontext.Fields{
		"performerID": performerID, "name": req.Name, "type": req.Type, "content": req.Content,
		"settings": req.Settings, "data": req.Data,
	})

	ctx.Logger().Text("get user QR codes quota")
	quota, isFreePlan, err := h.userHub.GetQRCodeQuota(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get user QR code quota", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count user total QR codes")
	totalQRCodes, err := h.qrCodeRepository.CountByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to count user total QR codes", err, appcontext.Fields{})
		return nil, err
	}

	if totalQRCodes >= quota {
		ctx.Logger().Error("user QR codes quota exceeded", err, appcontext.Fields{"quota": quota, "total": totalQRCodes})
		err = apperrors.User.FreePlanLimitReached
		if !isFreePlan {
			err = apperrors.User.ProPlanLimitReached
		}

		return nil, err
	}

	ctx.Logger().Text("create new QR code model")
	qrCode, err := domain.NewQRCode(performerID, req.Name, req.Type, req.Content, domain.QRCodeSettings{
		Color:    req.Settings.Color,
		HasLogo:  req.Settings.HasLogo,
		LogoData: req.Settings.LogoData,
		LogoName: req.Settings.LogoName,
		Style:    req.Settings.Style,
	}, req.Data)
	if err != nil {
		ctx.Logger().Error("failed to create new QR code model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist QR code in db")
	if err = h.qrCodeRepository.Create(ctx, *qrCode); err != nil {
		ctx.Logger().Error("failed to persist QR code in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create QR code request")
	return &dto.CreateQRCodeResponse{
		ID: qrCode.ID,
	}, nil
}
