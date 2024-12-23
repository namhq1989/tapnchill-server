package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mocktask "github.com/namhq1989/tapnchill-server/internal/mock/task"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createTaskTestSuite struct {
	suite.Suite
	handler            command.CreateTaskHandler
	mockCtrl           *gomock.Controller
	mockTaskRepository *mocktask.MockTaskRepository
	mockGoalRepository *mocktask.MockGoalRepository
	mockService        *mocktask.MockService
	mockUserHub        *mocktask.MockUserHub
}

func (s *createTaskTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createTaskTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockTaskRepository = mocktask.NewMockTaskRepository(s.mockCtrl)
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)
	s.mockService = mocktask.NewMockService(s.mockCtrl)
	s.mockUserHub = mocktask.NewMockUserHub(s.mockCtrl)

	s.handler = command.NewCreateTaskHandler(s.mockTaskRepository, s.mockGoalRepository, s.mockService, s.mockUserHub)
}

func (s *createTaskTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createTaskTestSuite) Test_1_Success() {
	// mock data
	var (
		goalID      = database.NewStringID()
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     goalID,
			UserID: performerID,
		}, nil)

	s.mockUserHub.EXPECT().
		GetTaskQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockTaskRepository.EXPECT().
		CountByGoalID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	s.mockTaskRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockGoalRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateTask(ctx, performerID, dto.CreateTaskRequest{
		GoalID:      goalID,
		Name:        "task name",
		Description: "task description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createTaskTestSuite) Test_2_Fail_InvalidName() {
	// call
	var (
		goalID      = database.NewStringID()
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     goalID,
			UserID: performerID,
		}, nil)

	s.mockUserHub.EXPECT().
		GetTaskQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockTaskRepository.EXPECT().
		CountByGoalID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateTask(ctx, performerID, dto.CreateTaskRequest{
		GoalID:      goalID,
		Name:        "",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *createTaskTestSuite) Test_2_Fail_InvalidGoalID() {
	// call
	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateTask(ctx, database.NewStringID(), dto.CreateTaskRequest{
		GoalID:      "",
		Name:        "task name",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *createTaskTestSuite) Test_2_Fail_NotGoalOwner() {
	// mock data
	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateTask(ctx, database.NewStringID(), dto.CreateTaskRequest{
		GoalID:      database.NewStringID(),
		Name:        "task name",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *createTaskTestSuite) Test_2_Fail_ResourceLimitReached() {
	// mock
	var (
		goalID      = database.NewStringID()
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     goalID,
			UserID: performerID,
		}, nil)

	s.mockUserHub.EXPECT().
		GetTaskQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockTaskRepository.EXPECT().
		CountByGoalID(gomock.Any(), gomock.Any()).
		Return(int64(10), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateTask(ctx, performerID, dto.CreateTaskRequest{
		GoalID:      goalID,
		Name:        "task name",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.FreePlanLimitReached, err)
}

//
// END OF CASES
//

func TestCreateTaskTestSuite(t *testing.T) {
	suite.Run(t, new(createTaskTestSuite))
}
