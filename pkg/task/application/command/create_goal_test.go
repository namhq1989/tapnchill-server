package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mocktask "github.com/namhq1989/tapnchill-server/internal/mock/task"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createGoalTestSuite struct {
	suite.Suite
	handler            command.CreateGoalHandler
	mockCtrl           *gomock.Controller
	mockGoalRepository *mocktask.MockGoalRepository
	mockUserHub        *mocktask.MockUserHub
}

func (s *createGoalTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createGoalTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)
	s.mockUserHub = mocktask.NewMockUserHub(s.mockCtrl)

	s.handler = command.NewCreateGoalHandler(s.mockGoalRepository, s.mockUserHub)
}

func (s *createGoalTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createGoalTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		GetGoalQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockGoalRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	s.mockGoalRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateGoal(ctx, database.NewStringID(), dto.CreateGoalRequest{
		Name:        "goal name",
		Description: "goal description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createGoalTestSuite) Test_2_Fail_InvalidName() {
	// mock
	s.mockUserHub.EXPECT().
		GetGoalQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockGoalRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateGoal(ctx, database.NewStringID(), dto.CreateGoalRequest{
		Name:        "",
		Description: "goal description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *createGoalTestSuite) Test_2_Fail_ResourceLimitReached() {
	// mock
	s.mockUserHub.EXPECT().
		GetGoalQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockGoalRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(10), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateGoal(ctx, database.NewStringID(), dto.CreateGoalRequest{
		Name:        "goal name",
		Description: "goal description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.FreePlanLimitReached, err)
}

//
// END OF CASES
//

func TestCreateGoalTestSuite(t *testing.T) {
	suite.Run(t, new(createGoalTestSuite))
}
