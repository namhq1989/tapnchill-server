package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockcommon "github.com/namhq1989/tapnchill-server/internal/mock/common"
	"github.com/namhq1989/tapnchill-server/pkg/common/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createFeedbackTestSuite struct {
	suite.Suite
	handler                command.CreateFeedbackHandler
	mockCtrl               *gomock.Controller
	mockFeedbackRepository *mockcommon.MockFeedbackRepository
}

func (s *createFeedbackTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createFeedbackTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockFeedbackRepository = mockcommon.NewMockFeedbackRepository(s.mockCtrl)

	s.handler = command.NewCreateFeedbackHandler(s.mockFeedbackRepository)
}

func (s *createFeedbackTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createFeedbackTestSuite) Test_1_Success() {
	// mock data
	s.mockFeedbackRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateFeedback(ctx, database.NewStringID(), dto.CreateFeedbackRequest{
		Email:    "test@gmail.com",
		Feedback: "feedback content",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createFeedbackTestSuite) Test_2_Fail_InvalidFeedbackContent() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateFeedback(ctx, database.NewStringID(), dto.CreateFeedbackRequest{
		Email:    "test@gmail.com",
		Feedback: "",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidFeedback, err)
}

//
// END OF CASES
//

func TestCreateFeedbackTestSuite(t *testing.T) {
	suite.Run(t, new(createFeedbackTestSuite))
}
