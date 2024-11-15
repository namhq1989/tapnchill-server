package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockhabit "github.com/namhq1989/tapnchill-server/internal/mock/habit"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type updateHabitTestSuite struct {
	suite.Suite
	handler             command.UpdateHabitHandler
	mockCtrl            *gomock.Controller
	mockHabitRepository *mockhabit.MockHabitRepository
	mockService         *mockhabit.MockService
}

func (s *updateHabitTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateHabitTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockHabitRepository = mockhabit.NewMockHabitRepository(s.mockCtrl)
	s.mockService = mockhabit.NewMockService(s.mockCtrl)

	s.handler = command.NewUpdateHabitHandler(s.mockHabitRepository, s.mockService)
}

func (s *updateHabitTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateHabitTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetHabitByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Habit{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockHabitRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateHabit(ctx, performerID, database.NewStringID(), dto.UpdateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DayOfWeeks: []int{1, 2, 3},
		Icon:       "icon.png",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateHabitTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetHabitByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateHabit(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DayOfWeeks: []int{1, 2, 3},
		Icon:       "icon.png",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *updateHabitTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetHabitByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateHabit(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DayOfWeeks: []int{1, 2, 3},
		Icon:       "icon.png",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestUpdateHabitTestSuite(t *testing.T) {
	suite.Run(t, new(updateHabitTestSuite))
}
