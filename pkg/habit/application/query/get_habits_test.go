package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockhabit "github.com/namhq1989/tapnchill-server/internal/mock/habit"
	"github.com/namhq1989/tapnchill-server/pkg/habit/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	"github.com/namhq1989/tapnchill-server/pkg/habit/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getHabitsTestSuite struct {
	suite.Suite
	handler     query.GetHabitsHandler
	mockCtrl    *gomock.Controller
	mockService *mockhabit.MockService
}

func (s *getHabitsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getHabitsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockhabit.NewMockService(s.mockCtrl)

	s.handler = query.NewGetHabitsHandler(s.mockService)
}

func (s *getHabitsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getHabitsTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetUserHabits(gomock.Any(), gomock.Any()).
		Return([]domain.Habit{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetHabits(ctx, performerID, dto.GetHabitsRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Habits))
}

//
// END OF CASES
//

func TestGetHabitsTestSuite(t *testing.T) {
	suite.Run(t, new(getHabitsTestSuite))
}
