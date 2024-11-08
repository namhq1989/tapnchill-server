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

type getTasksTestSuite struct {
	suite.Suite
	handler            query.GetTasksHandler
	mockCtrl           *gomock.Controller
	mockTaskRepository *mocktask.MockTaskRepository
}

func (s *getTasksTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getTasksTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockTaskRepository = mocktask.NewMockTaskRepository(s.mockCtrl)

	s.handler = query.NewGetTasksHandler(s.mockTaskRepository)
}

func (s *getTasksTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getTasksTestSuite) Test_1_Success() {
	// mock data
	s.mockTaskRepository.EXPECT().
		FindByFilter(gomock.Any(), gomock.Any()).
		Return([]domain.Task{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetTasks(ctx, database.NewStringID(), dto.GetTasksRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Tasks))
}

//
// END OF CASES
//

func TestGetTasksTestSuite(t *testing.T) {
	suite.Run(t, new(getTasksTestSuite))
}
