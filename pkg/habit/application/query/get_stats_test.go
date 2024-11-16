package query_test

import (
	"context"
	"testing"
	"time"

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

type getStatsTestSuite struct {
	suite.Suite
	handler     query.GetStatsHandler
	mockCtrl    *gomock.Controller
	mockService *mockhabit.MockService
}

func (s *getStatsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getStatsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockhabit.NewMockService(s.mockCtrl)

	s.handler = query.NewGetStatsHandler(s.mockService)
}

func (s *getStatsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getStatsTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetUserStats(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]domain.HabitDailyStats{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetStats(ctx, performerID, dto.GetStatsRequest{
		Date: time.Now().String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Stats))
}

//
// END OF CASES
//

func TestGetStatsTestSuite(t *testing.T) {
	suite.Run(t, new(getStatsTestSuite))
}
