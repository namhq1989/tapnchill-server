package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockhabit "github.com/namhq1989/tapnchill-server/internal/mock/habit"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type completeHabitTestSuite struct {
	suite.Suite
	handler                       command.CompleteHabitHandler
	mockCtrl                      *gomock.Controller
	mockHabitRepository           *mockhabit.MockHabitRepository
	mockHabitCompletionRepository *mockhabit.MockHabitCompletionRepository
	mockHabitDailyStatsRepository *mockhabit.MockHabitDailyStatsRepository
	mockCachingRepository         *mockhabit.MockCachingRepository
	mockService                   *mockhabit.MockService
}

func (s *completeHabitTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *completeHabitTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockHabitRepository = mockhabit.NewMockHabitRepository(s.mockCtrl)
	s.mockHabitCompletionRepository = mockhabit.NewMockHabitCompletionRepository(s.mockCtrl)
	s.mockHabitDailyStatsRepository = mockhabit.NewMockHabitDailyStatsRepository(s.mockCtrl)
	s.mockCachingRepository = mockhabit.NewMockCachingRepository(s.mockCtrl)
	s.mockService = mockhabit.NewMockService(s.mockCtrl)

	s.handler = command.NewCompleteHabitHandler(
		s.mockHabitRepository,
		s.mockHabitCompletionRepository,
		s.mockHabitDailyStatsRepository,
		s.mockCachingRepository,
		s.mockService,
	)
}

func (s *completeHabitTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *completeHabitTestSuite) Test_1_Success_FirstCompletion() {
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

	s.mockHabitDailyStatsRepository.
		EXPECT().
		FindByDate(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockHabitRepository.EXPECT().
		CountScheduledHabits(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockHabitDailyStatsRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitCompletionRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitDailyStatsRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserHabits(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserStats(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CompleteHabit(ctx, performerID, database.NewStringID(), dto.CompleteHabitRequest{
		Date: time.Now().Format(time.RFC3339),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *completeHabitTestSuite) Test_1_Success_NotFirstCompletion() {
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

	s.mockHabitDailyStatsRepository.
		EXPECT().
		FindByDate(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.HabitDailyStats{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	s.mockHabitCompletionRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitDailyStatsRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockHabitRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserHabits(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserStats(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CompleteHabit(ctx, performerID, database.NewStringID(), dto.CompleteHabitRequest{
		Date: time.Now().Format(time.RFC3339),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

//
// END OF CASES
//

func TestCompleteHabitTestSuite(t *testing.T) {
	suite.Run(t, new(completeHabitTestSuite))
}
