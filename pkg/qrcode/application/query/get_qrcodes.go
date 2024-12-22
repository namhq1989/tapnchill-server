package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/utils/pagetoken"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
)

type GetQRCodesHandler struct {
	qrCodeRepository domain.QRCodeRepository
}

func NewGetQRCodesHandler(qrCodeRepository domain.QRCodeRepository) GetQRCodesHandler {
	return GetQRCodesHandler{
		qrCodeRepository: qrCodeRepository,
	}
}

func (h GetQRCodesHandler) GetQRCodes(ctx *appcontext.AppContext, performerID string, req dto.GetQRCodesRequest) (*dto.GetQRCodesResponse, error) {
	ctx.Logger().Info("new get QR codes request", appcontext.Fields{
		"performerID": performerID, "pageToken": req.PageToken,
	})

	ctx.Logger().Text("create filter")
	filter, err := domain.NewQRCodeFilter(performerID, req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to create filter", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find QR codes in db")
	qrCodes, err := h.qrCodeRepository.FindByFilter(ctx, *filter)
	if err != nil {
		ctx.Logger().Error("failed to find QR codes in db", err, appcontext.Fields{})
		return nil, err
	}

	totalQRCodes := len(qrCodes)
	if totalQRCodes == 0 {
		ctx.Logger().Text("no QR codes found, respond")
		return &dto.GetQRCodesResponse{
			QRCodes:       make([]dto.QRCode, 0),
			NextPageToken: "",
		}, nil
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.QRCode, 0)
	for _, qrCode := range qrCodes {
		result = append(result, dto.QRCode{}.FromDomain(qrCode))
	}

	nextPageToken := ""
	if totalQRCodes == int(filter.Limit) {
		nextPageToken = pagetoken.NewWithTimestamp(qrCodes[totalQRCodes-1].CreatedAt)
	}

	ctx.Logger().Text("done get QR codes request")
	return &dto.GetQRCodesResponse{
		QRCodes:       result,
		NextPageToken: nextPageToken,
	}, nil
}
