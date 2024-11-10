package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mocktask "github.com/namhq1989/tapnchill-server/internal/mock/task"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/worker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createUserDefaultGoalTestSuite struct {
	suite.Suite
	handler            worker.CreateUserDefaultGoalHandler
	mockCtrl           *gomock.Controller
	mockGoalRepository *mocktask.MockGoalRepository
}

func (s *createUserDefaultGoalTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createUserDefaultGoalTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)

	s.handler = worker.NewCreateUserDefaultGoalHandler(s.mockGoalRepository)
}

func (s *createUserDefaultGoalTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createUserDefaultGoalTestSuite) Test_1_Success() {
	// mock data
	s.mockGoalRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.CreateUserDefaultGoal(ctx, domain.QueueCreateUserDefaultGoalPayload{
		UserID: database.NewStringID(),
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestCreateUserDefaultGoalTestSuite(t *testing.T) {
	suite.Run(t, new(createUserDefaultGoalTestSuite))
}
