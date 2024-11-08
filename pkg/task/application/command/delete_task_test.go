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

type deleteTaskTestSuite struct {
	suite.Suite
	handler            command.DeleteTaskHandler
	mockCtrl           *gomock.Controller
	mockTaskRepository *mocktask.MockTaskRepository
	mockGoalRepository *mocktask.MockGoalRepository
	mockService        *mocktask.MockService
}

func (s *deleteTaskTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *deleteTaskTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockTaskRepository = mocktask.NewMockTaskRepository(s.mockCtrl)
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)
	s.mockService = mocktask.NewMockService(s.mockCtrl)

	s.handler = command.NewDeleteTaskHandler(s.mockTaskRepository, s.mockGoalRepository, s.mockService)
}

func (s *deleteTaskTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *deleteTaskTestSuite) Test_1_Success() {
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
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockGoalRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteTask(ctx, performerID, database.NewStringID(), dto.DeleteTaskRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *deleteTaskTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteTask(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteTaskRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *deleteTaskTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteTask(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteTaskRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestDeleteTaskTestSuite(t *testing.T) {
	suite.Run(t, new(deleteTaskTestSuite))
}
