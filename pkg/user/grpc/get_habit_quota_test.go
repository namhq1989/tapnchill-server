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

type getHabitQuotaTestSuite struct {
	suite.Suite
	handler     grpc.GetHabitQuotaHandler
	mockCtrl    *gomock.Controller
	mockService *mockuser.MockService
}

func (s *getHabitQuotaTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getHabitQuotaTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockuser.NewMockService(s.mockCtrl)

	s.handler = grpc.NewGetHabitQuotaHandler(s.mockService)
}

func (s *getHabitQuotaTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getHabitQuotaTestSuite) Test_1_Success() {
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
	resp, err := s.handler.GetHabitQuota(ctx, &userpb.GetHabitQuotaRequest{
		TraceId: "trace-id",
		UserId:  database.NewStringID(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), domain.FreePlanMaxHabits, resp.GetLimit())
}

//
// END OF CASES
//

func TestGetHabitQuotaTestSuite(t *testing.T) {
	suite.Run(t, new(getHabitQuotaTestSuite))
}
