package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mocknote "github.com/namhq1989/tapnchill-server/internal/mock/note"
	"github.com/namhq1989/tapnchill-server/pkg/note/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type updateNoteTestSuite struct {
	suite.Suite
	handler            command.UpdateNoteHandler
	mockCtrl           *gomock.Controller
	mockNoteRepository *mocknote.MockNoteRepository
}

func (s *updateNoteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateNoteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockNoteRepository = mocknote.NewMockNoteRepository(s.mockCtrl)

	s.handler = command.NewUpdateNoteHandler(s.mockNoteRepository)
}

func (s *updateNoteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateNoteTestSuite) Test_1_Success() {
	// mock data
	var (
		performerID = database.NewStringID()
	)

	s.mockNoteRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Note{
			ID:     database.NewStringID(),
			UserID: performerID,
		}, nil)

	s.mockNoteRepository.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateNote(ctx, performerID, database.NewStringID(), dto.UpdateNoteRequest{
		Title:       "note title",
		Description: "note description",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateNoteTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockNoteRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateNote(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateNoteRequest{
		Title:       "note title",
		Description: "note description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *updateNoteTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockNoteRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Note{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateNote(ctx, database.NewStringID(), database.NewStringID(), dto.UpdateNoteRequest{
		Title:       "note title",
		Description: "note description",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestUpdateNoteTestSuite(t *testing.T) {
	suite.Run(t, new(updateNoteTestSuite))
}
