package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getPaymentCustomerPortalURLTestSuite struct {
	suite.Suite
	handler     query.GetPaymentCustomerPortalURLHandler
	mockCtrl    *gomock.Controller
	mockService *mockuser.MockService
}

func (s *getPaymentCustomerPortalURLTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getPaymentCustomerPortalURLTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockuser.NewMockService(s.mockCtrl)

	s.handler = query.NewGetPaymentCustomerPortalURLHandler(s.mockService)
}

func (s *getPaymentCustomerPortalURLTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getPaymentCustomerPortalURLTestSuite) Test_1_Success() {
	// mock data
	url := "https://google.com.vn"

	s.mockService.EXPECT().
		GetLemonsqueezyCustomerPortalURL(gomock.Any(), gomock.Any()).
		Return(&url, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetPaymentCustomerPortalURL(ctx, database.NewStringID(), dto.GetPaymentCustomerPortalURLRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

//
// END OF CASES
//

func TestGetPaymentCustomerPortalURLTestSuite(t *testing.T) {
	suite.Run(t, new(getPaymentCustomerPortalURLTestSuite))
}
