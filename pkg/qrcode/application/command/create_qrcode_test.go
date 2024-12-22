package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockqrcode "github.com/namhq1989/tapnchill-server/internal/mock/qrcode"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createQRCodeTestSuite struct {
	suite.Suite
	handler              command.CreateQRCodeHandler
	mockCtrl             *gomock.Controller
	mockQRCodeRepository *mockqrcode.MockQRCodeRepository
	mockUserHub          *mockqrcode.MockUserHub
}

func (s *createQRCodeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createQRCodeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQRCodeRepository = mockqrcode.NewMockQRCodeRepository(s.mockCtrl)
	s.mockUserHub = mockqrcode.NewMockUserHub(s.mockCtrl)

	s.handler = command.NewCreateQRCodeHandler(s.mockQRCodeRepository, s.mockUserHub)
}

func (s *createQRCodeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createQRCodeTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		GetQRCodeQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockQRCodeRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	s.mockQRCodeRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateQRCode(ctx, database.NewStringID(), dto.CreateQRCodeRequest{
		Name:     "qrcode name",
		Type:     "qrcode type",
		Content:  "qrcode content",
		Settings: dto.QRCodeSettings{},
		Data:     "",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createQRCodeTestSuite) Test_2_Fail_InvalidName() {
	// mock
	s.mockUserHub.EXPECT().
		GetQRCodeQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockQRCodeRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateQRCode(ctx, database.NewStringID(), dto.CreateQRCodeRequest{
		Name:     "",
		Type:     "qrcode type",
		Content:  "qrcode content",
		Settings: dto.QRCodeSettings{},
		Data:     "",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *createQRCodeTestSuite) Test_2_Fail_ResourceLimitReached() {
	// mock
	s.mockUserHub.EXPECT().
		GetQRCodeQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockQRCodeRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(10), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateQRCode(ctx, database.NewStringID(), dto.CreateQRCodeRequest{
		Name:     "qrcode name",
		Type:     "qrcode type",
		Content:  "qrcode content",
		Settings: dto.QRCodeSettings{},
		Data:     "",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.ResourceLimitReached, err)
}

//
// END OF CASES
//

func TestCreateQRCodeTestSuite(t *testing.T) {
	suite.Run(t, new(createQRCodeTestSuite))
}
