package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockhabit "github.com/namhq1989/tapnchill-server/internal/mock/habit"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type changeHabitStatusTestSuite struct {
	suite.Suite
	handler                       command.ChangeHabitStatusHandler
	mockCtrl                      *gomock.Controller
	mockHabitRepository           *mockhabit.MockHabitRepository
	mockHabitDailyStatsRepository *mockhabit.MockHabitDailyStatsRepository
	mockService                   *mockhabit.MockService
}

func (s *changeHabitStatusTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *changeHabitStatusTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockHabitRepository = mockhabit.NewMockHabitRepository(s.mockCtrl)
	s.mockService = mockhabit.NewMockService(s.mockCtrl)

	s.handler = command.NewChangeHabitStatusHandler(s.mockHabitRepository, s.mockHabitDailyStatsRepository, s.mockService)
}

func (s *changeHabitStatusTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *changeHabitStatusTestSuite) Test_1_Success() {
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

	s.mockService.EXPECT().
		DeleteUserCaching(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitDailyStatsRepository.EXPECT().
		FindByDate(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.HabitDailyStats{
			ID:           database.NewStringID(),
			ScheduledIDs: make([]string, 0),
		}, nil)

	s.mockHabitDailyStatsRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeHabitStatus(ctx, performerID, database.NewStringID(), dto.ChangeHabitStatusRequest{
		Status: domain.HabitStatusActive.String(),
		Date:   time.Now().Format(time.RFC3339),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *changeHabitStatusTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockService.EXPECT().
		GetHabitByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeHabitStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeHabitStatusRequest{
		Status: domain.HabitStatusActive.String(),
		Date:   time.Now().Format(time.RFC3339),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *changeHabitStatusTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockService.EXPECT().
		GetHabitByID(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeHabitStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeHabitStatusRequest{
		Status: domain.HabitStatusActive.String(),
		Date:   time.Now().Format(time.RFC3339),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *changeHabitStatusTestSuite) Test_2_Fail_InvalidStatus() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeHabitStatus(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeHabitStatusRequest{
		Status: "invalid status",
		Date:   time.Now().Format(time.RFC3339),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

//
// END OF CASES
//

func TestChangeHabitStatusTestSuite(t *testing.T) {
	suite.Run(t, new(changeHabitStatusTestSuite))
}
