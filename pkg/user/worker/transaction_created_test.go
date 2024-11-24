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

type transactionCompletedTestSuite struct {
	suite.Suite
	handler                           worker.TransactionCompletedHandler
	mockCtrl                          *gomock.Controller
	mockUserRepository                *mockuser.MockUserRepository
	mockSubscriptionHistoryRepository *mockuser.MockSubscriptionHistoryRepository
	mockCachingRepository             *mockuser.MockCachingRepository
}

func (s *transactionCompletedTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *transactionCompletedTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockSubscriptionHistoryRepository = mockuser.NewMockSubscriptionHistoryRepository(s.mockCtrl)
	s.mockCachingRepository = mockuser.NewMockCachingRepository(s.mockCtrl)

	s.handler = worker.NewTransactionCompletedHandler(s.mockUserRepository, s.mockSubscriptionHistoryRepository, s.mockCachingRepository)
}

func (s *transactionCompletedTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *transactionCompletedTestSuite) Test_1_Success() {
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
	err := s.handler.TransactionCompleted(ctx, domain.QueueTransactionCompletedPayload{
		UserID:         database.NewStringID(),
		SubscriptionID: "subscription-id",
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestTransactionCompletedTestSuite(t *testing.T) {
	suite.Run(t, new(transactionCompletedTestSuite))
}
