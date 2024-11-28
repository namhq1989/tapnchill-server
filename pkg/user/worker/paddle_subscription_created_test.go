package worker_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/worker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type paddleSubscriptionCreatedTestSuite struct {
	suite.Suite
	handler                           worker.PaddleSubscriptionCreatedHandler
	mockCtrl                          *gomock.Controller
	mockUserRepository                *mockuser.MockUserRepository
	mockSubscriptionHistoryRepository *mockuser.MockSubscriptionHistoryRepository
}

func (s *paddleSubscriptionCreatedTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *paddleSubscriptionCreatedTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockSubscriptionHistoryRepository = mockuser.NewMockSubscriptionHistoryRepository(s.mockCtrl)

	s.handler = worker.NewPaddleSubscriptionCreatedHandler(s.mockUserRepository, s.mockSubscriptionHistoryRepository)
}

func (s *paddleSubscriptionCreatedTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *paddleSubscriptionCreatedTestSuite) Test_1_Success() {
	// mock data
	s.mockUserRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID: database.NewStringID(),
		}, nil)

	s.mockSubscriptionHistoryRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockUserRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewWorker(context.Background())
	err := s.handler.PaddleSubscriptionCreated(ctx, domain.QueuePaddleSubscriptionCreatedPayload{
		UserID:         database.NewStringID(),
		SubscriptionID: "subscription-id",
		NextBilledAt:   time.Now(),
		CustomerID:     "customer-id",
		Items:          []string{"item-1", "item-2"},
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestPaddleSubscriptionCreatedTestSuite(t *testing.T) {
	suite.Run(t, new(paddleSubscriptionCreatedTestSuite))
}
