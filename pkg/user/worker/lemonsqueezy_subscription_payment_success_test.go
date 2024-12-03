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

type lemonsqueezySubscriptionPaymentSuccessTestSuite struct {
	suite.Suite
	handler                           worker.LemonsqueezySubscriptionPaymentSuccessHandler
	mockCtrl                          *gomock.Controller
	mockUserRepository                *mockuser.MockUserRepository
	mockSubscriptionHistoryRepository *mockuser.MockSubscriptionHistoryRepository
	mockCachingRepository             *mockuser.MockCachingRepository
	mockExternalAPIRepository         *mockuser.MockExternalAPIRepository
}

func (s *lemonsqueezySubscriptionPaymentSuccessTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *lemonsqueezySubscriptionPaymentSuccessTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockSubscriptionHistoryRepository = mockuser.NewMockSubscriptionHistoryRepository(s.mockCtrl)
	s.mockCachingRepository = mockuser.NewMockCachingRepository(s.mockCtrl)
	s.mockExternalAPIRepository = mockuser.NewMockExternalAPIRepository(s.mockCtrl)

	s.handler = worker.NewLemonsqueezySubscriptionPaymentSuccessHandler(
		s.mockUserRepository,
		s.mockSubscriptionHistoryRepository,
		s.mockCachingRepository,
		s.mockExternalAPIRepository,
	)
}

func (s *lemonsqueezySubscriptionPaymentSuccessTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *lemonsqueezySubscriptionPaymentSuccessTestSuite) Test_1_Success() {
	// mock data
	renewsAt := time.Now().AddDate(1, 0, 0)

	s.mockUserRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID: database.NewStringID(),
		}, nil)

	s.mockExternalAPIRepository.EXPECT().
		GetLemonsqueezyInvoiceData(gomock.Any(), gomock.Any()).
		Return(&domain.GetLemonsqueezyInvoiceDataResult{
			SubscriptionID: "subscription-id",
			CustomerID:     "customer-id",
			VariantID:      "variant-id",
			RenewsAt:       &renewsAt,
		}, nil)

	s.mockSubscriptionHistoryRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockUserRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserByID(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewWorker(context.Background())
	err := s.handler.LemonsqueezySubscriptionPaymentSuccess(ctx, domain.QueueLemonsqueezySubscriptionPaymentSuccessPayload{
		UserID:    database.NewStringID(),
		InvoiceID: "invoice-id",
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestLemonsqueezySubscriptionPaymentSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(lemonsqueezySubscriptionPaymentSuccessTestSuite))
}
