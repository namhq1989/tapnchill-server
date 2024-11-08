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

type updateGoalTestSuite struct {
	suite.Suite
	handler            command.UpdateGoalHandler
	mockCtrl           *gomock.Controller
	mockGoalRepository *mocktask.MockGoalRepository
	mockService        *mocktask.MockService
}

func (s *updateGoalTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateGoalTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)
	s.mockService = mocktask.NewMockService(s.mockCtrl)

	s.handler = command.NewUpdateGoalHandler(s.mockGoalRepository, s.mockService)
}

func (s *updateGoalTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateGoalTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockGoalRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateGoal(ctx, performerID, database.NewStringID(), dto.UpdateGoalRequest{
		Name:        "goal name",
		Description: "goal description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateGoalTestSuite) Test_2_Fail_InvalidName() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Goal{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateGoal(ctx, performerID, database.NewStringID(), dto.UpdateGoalRequest{
		Name:        "",
		Description: "goal description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *updateGoalTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateGoal(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateGoalRequest{
		Name:        "goal name",
		Description: "goal description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *updateGoalTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetGoalByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateGoal(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateGoalRequest{
		Name:        "goal name",
		Description: "goal description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestUpdateGoalTestSuite(t *testing.T) {
	suite.Run(t, new(updateGoalTestSuite))
}
