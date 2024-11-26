package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mocknote "github.com/namhq1989/tapnchill-server/internal/mock/note"
	"github.com/namhq1989/tapnchill-server/pkg/note/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createNoteTestSuite struct {
	suite.Suite
	handler            command.CreateNoteHandler
	mockCtrl           *gomock.Controller
	mockNoteRepository *mocknote.MockNoteRepository
	mockUserHub        *mocknote.MockUserHub
}

func (s *createNoteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createNoteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockNoteRepository = mocknote.NewMockNoteRepository(s.mockCtrl)
	s.mockUserHub = mocknote.NewMockUserHub(s.mockCtrl)

	s.handler = command.NewCreateNoteHandler(s.mockNoteRepository, s.mockUserHub)
}

func (s *createNoteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createNoteTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		GetNoteQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockNoteRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	s.mockNoteRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateNote(ctx, database.NewStringID(), dto.CreateNoteRequest{
		Title:       "note title",
		Description: "note description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createNoteTestSuite) Test_2_Fail_InvalidName() {
	// mock
	s.mockUserHub.EXPECT().
		GetNoteQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockNoteRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateNote(ctx, database.NewStringID(), dto.CreateNoteRequest{
		Title:       "",
		Description: "note description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

func (s *createNoteTestSuite) Test_2_Fail_ResourceLimitReached() {
	// mock
	s.mockUserHub.EXPECT().
		GetNoteQuota(gomock.Any(), gomock.Any()).
		Return(int64(5), nil)

	s.mockNoteRepository.EXPECT().
		CountByUserID(gomock.Any(), gomock.Any()).
		Return(int64(10), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateNote(ctx, database.NewStringID(), dto.CreateNoteRequest{
		Title:       "note title",
		Description: "note description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.ResourceLimitReached, err)
}

//
// END OF CASES
//

func TestCreateNoteTestSuite(t *testing.T) {
	suite.Run(t, new(createNoteTestSuite))
}
