package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getMeTestSuite struct {
	suite.Suite
	handler     query.GetMeHandler
	mockCtrl    *gomock.Controller
	mockService *mockuser.MockService
}

func (s *getMeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getMeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockuser.NewMockService(s.mockCtrl)

	s.handler = query.NewGetMeHandler(s.mockService)
}

func (s *getMeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getMeTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		GetUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetMe(ctx, database.NewStringID(), dto.GetMeRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

//
// END OF CASES
//

func TestGetMeTestSuite(t *testing.T) {
	suite.Run(t, new(getMeTestSuite))
}
