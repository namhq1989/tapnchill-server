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

type getTaskQuotaTestSuite struct {
	suite.Suite
	handler     grpc.GetTaskQuotaHandler
	mockCtrl    *gomock.Controller
	mockService *mockuser.MockService
}

func (s *getTaskQuotaTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getTaskQuotaTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockuser.NewMockService(s.mockCtrl)

	s.handler = grpc.NewGetTaskQuotaHandler(s.mockService)
}

func (s *getTaskQuotaTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getTaskQuotaTestSuite) Test_1_Success() {
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
	resp, err := s.handler.GetTaskQuota(ctx, &userpb.GetTaskQuotaRequest{
		TraceId: "trace-id",
		UserId:  database.NewStringID(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), domain.FreePlanMaxTaskPerGoal, resp.GetLimit())
}

//
// END OF CASES
//

func TestGetTaskQuotaTestSuite(t *testing.T) {
	suite.Run(t, new(getTaskQuotaTestSuite))
}
