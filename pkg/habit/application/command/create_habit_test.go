package command_test

import (
	"context"
	"testing"

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

type createHabitTestSuite struct {
	suite.Suite
	handler               command.CreateHabitHandler
	mockCtrl              *gomock.Controller
	mockHabitRepository   *mockhabit.MockHabitRepository
	mockCachingRepository *mockhabit.MockCachingRepository
}

func (s *createHabitTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createHabitTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockHabitRepository = mockhabit.NewMockHabitRepository(s.mockCtrl)
	s.mockCachingRepository = mockhabit.NewMockCachingRepository(s.mockCtrl)

	s.handler = command.NewCreateHabitHandler(s.mockHabitRepository, s.mockCachingRepository)
}

func (s *createHabitTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createHabitTestSuite) Test_1_Success() {
	// mock data
	s.mockHabitRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockCachingRepository.EXPECT().
		DeleteUserHabits(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateHabit(ctx, database.NewStringID(), dto.CreateHabitRequest{
		Name:       "habit name",
		Goal:       "habit goal",
		DaysOfWeek: []int{1, 2, 3},
		Icon:       "icon.png",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createHabitTestSuite) Test_2_Fail_InvalidName() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateHabit(ctx, database.NewStringID(), dto.CreateHabitRequest{
		Name:       "",
		Goal:       "habit goal",
		DaysOfWeek: []int{1, 2, 3},
		Icon:       "icon.png",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

//
// END OF CASES
//

func TestCreateHabitTestSuite(t *testing.T) {
	suite.Run(t, new(createHabitTestSuite))
}
