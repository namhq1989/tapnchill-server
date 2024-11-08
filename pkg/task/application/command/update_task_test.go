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

type updateTaskTestSuite struct {
	suite.Suite
	handler            command.UpdateTaskHandler
	mockCtrl           *gomock.Controller
	mockTaskRepository *mocktask.MockTaskRepository
	mockService        *mocktask.MockService
}

func (s *updateTaskTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateTaskTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockTaskRepository = mocktask.NewMockTaskRepository(s.mockCtrl)
	s.mockService = mocktask.NewMockService(s.mockCtrl)

	s.handler = command.NewUpdateTaskHandler(s.mockTaskRepository, s.mockService)
}

func (s *updateTaskTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateTaskTestSuite) Test_1_Success() {
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

	s.mockTaskRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateTask(ctx, performerID, database.NewStringID(), dto.UpdateTaskRequest{
		Name:        "task name",
		Description: "task description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateTaskTestSuite) Test_2_Fail_InvalidName() {
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

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateTask(ctx, performerID, database.NewStringID(), dto.UpdateTaskRequest{
		Name:        "",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *updateTaskTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateTask(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateTaskRequest{
		Name:        "task name",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *updateTaskTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetTaskByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateTask(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateTaskRequest{
		Name:        "task name",
		Description: "task description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestUpdateTaskTestSuite(t *testing.T) {
	suite.Run(t, new(updateTaskTestSuite))
}
