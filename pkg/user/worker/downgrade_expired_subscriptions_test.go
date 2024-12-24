package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/worker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type downgradeExpiredSubscriptionsTestSuite struct {
	suite.Suite
	handler            worker.DowngradeExpiredSubscriptionsHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockuser.MockUserRepository
}

func (s *downgradeExpiredSubscriptionsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *downgradeExpiredSubscriptionsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)

	s.handler = worker.NewDowngradeExpiredSubscriptionsHandler(
		s.mockUserRepository,
	)
}

func (s *downgradeExpiredSubscriptionsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *downgradeExpiredSubscriptionsTestSuite) Test_1_Success() {
	// mock data
	s.mockUserRepository.EXPECT().
		DowngradeAllExpiredSubscriptions(gomock.Any()).
		Return(int64(1), nil)

	// call
	ctx := appcontext.NewWorker(context.Background())
	err := s.handler.DowngradeExpiredSubscriptions(ctx, domain.QueueDowngradeExpiredSubscriptionsPayload{})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestDowngradeExpiredSubscriptionsTestSuite(t *testing.T) {
	suite.Run(t, new(downgradeExpiredSubscriptionsTestSuite))
}
