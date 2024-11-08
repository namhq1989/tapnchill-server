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

type deleteGoalTestSuite struct {
	suite.Suite
	handler            command.DeleteGoalHandler
	mockCtrl           *gomock.Controller
	mockGoalRepository *mocktask.MockGoalRepository
}

func (s *deleteGoalTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *deleteGoalTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)

	s.handler = command.NewDeleteGoalHandler(s.mockGoalRepository)
}

func (s *deleteGoalTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *deleteGoalTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockGoalRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockGoalRepository.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteGoal(ctx, performerID, database.NewStringID(), dto.DeleteGoalRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *deleteGoalTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockGoalRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteGoal(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteGoalRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *deleteGoalTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockGoalRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteGoal(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteGoalRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *deleteGoalTestSuite) Test_2_Fail_StillHasTasksRemaining() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockGoalRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: performerID,
			Stats: domain.GoalStats{
				TotalTask: 2,
			},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.DeleteGoal(ctx, performerID, database.NewStringID(), dto.DeleteGoalRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Task.GoalDeleteErrorTasksRemaining, err)
}

//
// END OF CASES
//

func TestDeleteGoalTestSuite(t *testing.T) {
	suite.Run(t, new(deleteGoalTestSuite))
}
