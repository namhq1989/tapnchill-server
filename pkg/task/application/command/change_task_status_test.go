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

type ChangeTaskStatusTestSuite struct {
	suite.Suite
	handler            command.ChangeTaskStatusHandler
	mockCtrl           *gomock.Controller
	mockTaskRepository *mocktask.MockTaskRepository
	mockGoalRepository *mocktask.MockGoalRepository
	mockService        *mocktask.MockService
}

func (s *ChangeTaskStatusTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *ChangeTaskStatusTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockTaskRepository = mocktask.NewMockTaskRepository(s.mockCtrl)
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)
	s.mockService = mocktask.NewMockService(s.mockCtrl)

	s.handler = command.NewChangeTaskStatusHandler(s.mockTaskRepository, s.mockGoalRepository, s.mockService)
}

func (s *ChangeTaskStatusTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *ChangeTaskStatusTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Task{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockTaskRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockGoalRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeTaskStatus(ctx, performerID, database.NewStringID(), dto.ChangeTaskStatusRequest{
		Status: domain.TaskStatusDone.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *ChangeTaskStatusTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeTaskStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeTaskStatusRequest{
		Status: domain.TaskStatusDone.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *ChangeTaskStatusTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeTaskStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeTaskStatusRequest{
		Status: domain.TaskStatusDone.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *ChangeTaskStatusTestSuite) Test_2_Fail_InvalidStatus() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeTaskStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeTaskStatusRequest{
		Status: "invalid status",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

//
// END OF CASES
//

func TestChangeTaskStatusTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeTaskStatusTestSuite))
}
