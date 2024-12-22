package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockqrcode "github.com/namhq1989/tapnchill-server/internal/mock/qrcode"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/domain"
	"github.com/namhq1989/tapnchill-server/pkg/qrcode/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getQRCodesTestSuite struct {
	suite.Suite
	handler              query.GetQRCodesHandler
	mockCtrl             *gomock.Controller
	mockQRCodeRepository *mockqrcode.MockQRCodeRepository
}

func (s *getQRCodesTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getQRCodesTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQRCodeRepository = mockqrcode.NewMockQRCodeRepository(s.mockCtrl)

	s.handler = query.NewGetQRCodesHandler(s.mockQRCodeRepository)
}

func (s *getQRCodesTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getQRCodesTestSuite) Test_1_Success() {
	// mock data
	s.mockQRCodeRepository.EXPECT().
		FindByFilter(gomock.Any(), gomock.Any()).
		Return([]domain.QRCode{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetQRCodes(ctx, database.NewStringID(), dto.GetQRCodesRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.QRCodes))
}

//
// END OF CASES
//

func TestGetQRCodesTestSuite(t *testing.T) {
	suite.Run(t, new(getQRCodesTestSuite))
}
