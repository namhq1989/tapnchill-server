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

type createHabitTestSuite struct {
	suite.Suite
	handler                       command.CreateHabitHandler
	mockCtrl                      *gomock.Controller
	mockHabitRepository           *mockhabit.MockHabitRepository
	mockHabitDailyStatsRepository *mockhabit.MockHabitDailyStatsRepository
	mockService                   *mockhabit.MockService
	mockUserHub                   *mockhabit.MockUserHub
}

func (s *createHabitTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createHabitTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockHabitRepository = mockhabit.NewMockHabitRepository(s.mockCtrl)
	s.mockHabitDailyStatsRepository = mockhabit.NewMockHabitDailyStatsRepository(s.mockCtrl)
	s.mockService = mockhabit.NewMockService(s.mockCtrl)
	s.mockUserHub = mockhabit.NewMockUserHub(s.mockCtrl)

	s.handler = command.NewCreateHabitHandler(s.mockHabitRepository, s.mockHabitDailyStatsRepository, s.mockService, s.mockUserHub)
}

func (s *createHabitTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createHabitTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		GetHabitQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockHabitRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	s.mockHabitRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
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
	resp, err := s.handler.CreateHabit(ctx, database.NewStringID(), dto.CreateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DaysOfWeek: []int{1, 2, 3},
		Icon:       "icon.png",
		Date:       time.Now().Format(time.RFC3339),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createHabitTestSuite) Test_2_Fail_InvalidName() {
	// mock
	s.mockUserHub.EXPECT().
		GetHabitQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockHabitRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateHabit(ctx, database.NewStringID(), dto.CreateHabitRequest{
		Name:       "",
		Goal:       "habit goal",
		DaysOfWeek: []int{1, 2, 3},
		Icon:       "icon.png",
		Date:       time.Now().Format(time.RFC3339),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *createHabitTestSuite) Test_2_Fail_ResourceLimitReached() {
	// mock
	s.mockUserHub.EXPECT().
		GetHabitQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), true, nil)

	s.mockHabitRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(10), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateHabit(ctx, database.NewStringID(), dto.CreateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DaysOfWeek: []int{1, 2, 3},
		Icon:       "icon.png",
		Date:       time.Now().Format(time.RFC3339),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.FreePlanLimitReached, err)
}

//
// END OF CASES
//

func TestCreateHabitTestSuite(t *testing.T) {
	suite.Run(t, new(createHabitTestSuite))
}
