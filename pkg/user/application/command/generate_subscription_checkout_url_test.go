package command_test

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type generateSubscriptionCheckoutURLTestSuite struct {
	suite.Suite
	handler                   command.GenerateSubscriptionCheckoutURLHandler
	mockCtrl                  *gomock.Controller
	mockExternalAPIRepository *mockuser.MockExternalAPIRepository
}

func (s *generateSubscriptionCheckoutURLTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *generateSubscriptionCheckoutURLTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExternalAPIRepository = mockuser.NewMockExternalAPIRepository(s.mockCtrl)

	s.handler = command.NewGenerateSubscriptionCheckoutURLHandler(s.mockExternalAPIRepository)
}

func (s *generateSubscriptionCheckoutURLTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *generateSubscriptionCheckoutURLTestSuite) Test_1_Success_Monthly() {
	// mock data
	var url = "https://valid-url"

	s.mockExternalAPIRepository.EXPECT().
		GenerateLemonsqueezyCheckoutURL(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&url, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GenerateSubscriptionCheckoutURL(ctx, database.NewStringID(), dto.GenerateSubscriptionCheckoutURLRequest{
		SubscriptionID: domain.SubscriptionIDMonthly,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), base64.StdEncoding.EncodeToString([]byte(url)), resp.CheckoutURL)
}

func (s *generateSubscriptionCheckoutURLTestSuite) Test_1_Success_Yearly() {
	// mock data
	var url = "https://valid-url"

	s.mockExternalAPIRepository.EXPECT().
		GenerateLemonsqueezyCheckoutURL(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&url, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GenerateSubscriptionCheckoutURL(ctx, database.NewStringID(), dto.GenerateSubscriptionCheckoutURLRequest{
		SubscriptionID: domain.SubscriptionIDYearly,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), base64.StdEncoding.EncodeToString([]byte(url)), resp.CheckoutURL)
}

func (s *generateSubscriptionCheckoutURLTestSuite) Test_2_Fail_InvalidSubscriptionID() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GenerateSubscriptionCheckoutURL(ctx, database.NewStringID(), dto.GenerateSubscriptionCheckoutURLRequest{
		SubscriptionID: "invalid-id",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

//
// END OF CASES
//

func TestGenerateSubscriptionCheckoutURLTestSuite(t *testing.T) {
	suite.Run(t, new(generateSubscriptionCheckoutURLTestSuite))
}
