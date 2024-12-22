package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockqrcode "github.com/namhq1989/tapnchill-server/internal/mock/qrcode"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type updateQRCodeTestSuite struct {
	suite.Suite
	handler              command.UpdateQRCodeHandler
	mockCtrl             *gomock.Controller
	mockQRCodeRepository *mockqrcode.MockQRCodeRepository
}

func (s *updateQRCodeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateQRCodeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQRCodeRepository = mockqrcode.NewMockQRCodeRepository(s.mockCtrl)

	s.handler = command.NewUpdateQRCodeHandler(s.mockQRCodeRepository)
}

func (s *updateQRCodeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateQRCodeTestSuite) Test_1_Success() {
	// mock data
	var (
		qrCodeID    = database.NewStringID()
		performerID = database.NewStringID()
	)

	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.QRCode{
			ID:     qrCodeID,
			UserID: performerID,
		}, nil)

	s.mockQRCodeRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateQRCode(ctx, performerID, qrCodeID, dto.UpdateQRCodeRequest{
		Name: "qrcode name",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateQRCodeTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.UpdateQRCode(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateQRCodeRequest{
		Name: "qrcode name",
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *updateQRCodeTestSuite) Test_2_Fail_InvalidName() {
	// mock
	var (
		qrCodeID    = database.NewStringID()
		performerID = database.NewStringID()
	)

	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.QRCode{
			ID:     qrCodeID,
			UserID: performerID,
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateQRCode(ctx, performerID, qrCodeID, dto.UpdateQRCodeRequest{
		Name: "",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *updateQRCodeTestSuite) Test_2_Fail_NotOwner() {
	// mock
	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.QRCode{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateQRCode(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateQRCodeRequest{
		Name: "qrcode name",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestupdateQRCodeTestSuite(t *testing.T) {
	suite.Run(t, new(updateQRCodeTestSuite))
}
