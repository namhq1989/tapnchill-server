package grpc_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	"github.com/namhq1989/tapnchill-server/internal/genproto/userpb"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getQRCodeQuotaTestSuite struct {
	suite.Suite
	handler     grpc.GetQRCodeQuotaHandler
	mockCtrl    *gomock.Controller
	mockService *mockuser.MockService
}

func (s *getQRCodeQuotaTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getQRCodeQuotaTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockuser.NewMockService(s.mockCtrl)

	s.handler = grpc.NewGetQRCodeQuotaHandler(s.mockService)
}

func (s *getQRCodeQuotaTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getQRCodeQuotaTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		GetUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID: database.NewStringID(),
			Subscription: domain.UserSubscription{
				Plan: "free",
			},
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetQRCodeQuota(ctx, &userpb.GetQRCodeQuotaRequest{
		TraceId: "trace-id",
		UserId:  database.NewStringID(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), domain.FreePlanMaxQRCodes, resp.GetLimit())
}

//
// END OF CASES
//

func TestGetQRCodeQuotaTestSuite(t *testing.T) {
	suite.Run(t, new(getQRCodeQuotaTestSuite))
}
