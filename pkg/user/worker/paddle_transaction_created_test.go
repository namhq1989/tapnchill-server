package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/worker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type paddleTransactionCompletedTestSuite struct {
	suite.Suite
	handler                           worker.PaddleTransactionCompletedHandler
	mockCtrl                          *gomock.Controller
	mockUserRepository                *mockuser.MockUserRepository
	mockSubscriptionHistoryRepository *mockuser.MockSubscriptionHistoryRepository
	mockCachingRepository             *mockuser.MockCachingRepository
}

func (s *paddleTransactionCompletedTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *paddleTransactionCompletedTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockSubscriptionHistoryRepository = mockuser.NewMockSubscriptionHistoryRepository(s.mockCtrl)
	s.mockCachingRepository = mockuser.NewMockCachingRepository(s.mockCtrl)

	s.handler = worker.NewPaddleTransactionCompletedHandler(s.mockUserRepository, s.mockSubscriptionHistoryRepository, s.mockCachingRepository)
}

func (s *paddleTransactionCompletedTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *paddleTransactionCompletedTestSuite) Test_1_Success() {
	// mock data
	s.mockUserRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID: database.NewStringID(),
		}, nil)

	s.mockSubscriptionHistoryRepository.EXPECT().
		FindBySourceID(gomock.Any(), gomock.Any()).
		Return(&domain.SubscriptionHistory{
			ID: database.NewStringID(),
		}, nil)

	s.mockSubscriptionHistoryRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockUserRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserByID(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewWorker(context.Background())
	err := s.handler.PaddleTransactionCompleted(ctx, domain.QueuePaddleTransactionCompletedPayload{
		UserID:         database.NewStringID(),
		SubscriptionID: "subscription-id",
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestPaddleTransactionCompletedTestSuite(t *testing.T) {
	suite.Run(t, new(paddleTransactionCompletedTestSuite))
}
