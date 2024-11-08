package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mocktask "github.com/namhq1989/tapnchill-server/internal/mock/task"
	"github.com/namhq1989/tapnchill-server/pkg/task/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/task/domain"
	"github.com/namhq1989/tapnchill-server/pkg/task/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getGoalsTestSuite struct {
	suite.Suite
	handler            query.GetGoalsHandler
	mockCtrl           *gomock.Controller
	mockGoalRepository *mocktask.MockGoalRepository
}

func (s *getGoalsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getGoalsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGoalRepository = mocktask.NewMockGoalRepository(s.mockCtrl)

	s.handler = query.NewGetGoalsHandler(s.mockGoalRepository)
}

func (s *getGoalsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getGoalsTestSuite) Test_1_Success() {
	// mock data
	s.mockGoalRepository.EXPECT().
		FindByFilter(gomock.Any(), gomock.Any()).
		Return([]domain.Goal{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetGoals(ctx, database.NewStringID(), dto.GetGoalsRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Goals))
}

//
// END OF CASES
//

func TestGetGoalsTestSuite(t *testing.T) {
	suite.Run(t, new(getGoalsTestSuite))
}
