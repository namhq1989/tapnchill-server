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

type deleteQRCodeTestSuite struct {
	suite.Suite
	handler              command.DeleteQRCodeHandler
	mockCtrl             *gomock.Controller
	mockQRCodeRepository *mockqrcode.MockQRCodeRepository
}

func (s *deleteQRCodeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *deleteQRCodeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQRCodeRepository = mockqrcode.NewMockQRCodeRepository(s.mockCtrl)

	s.handler = command.NewDeleteQRCodeHandler(s.mockQRCodeRepository)
}

func (s *deleteQRCodeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *deleteQRCodeTestSuite) Test_1_Success() {
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
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteQRCode(ctx, performerID, qrCodeID, dto.DeleteQRCodeRequest{})

	assert.Nil(s.T(), err)
}

func (s *deleteQRCodeTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteQRCode(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteQRCodeRequest{})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *deleteQRCodeTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockQRCodeRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.QRCode{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteQRCode(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteQRCodeRequest{})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestDeleteQRCodeTestSuite(t *testing.T) {
	suite.Run(t, new(deleteQRCodeTestSuite))
}
