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

type deleteNoteTestSuite struct {
	suite.Suite
	handler            command.DeleteNoteHandler
	mockCtrl           *gomock.Controller
	mockNoteRepository *mocknote.MockNoteRepository
}

func (s *deleteNoteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *deleteNoteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockNoteRepository = mocknote.NewMockNoteRepository(s.mockCtrl)

	s.handler = command.NewDeleteNoteHandler(s.mockNoteRepository)
}

func (s *deleteNoteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *deleteNoteTestSuite) Test_1_Success() {
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
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteNote(ctx, performerID, database.NewStringID(), dto.DeleteNoteRequest{})

	assert.Nil(s.T(), err)
}

func (s *deleteNoteTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockNoteRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.NotFound)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteNote(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteNoteRequest{})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

func (s *deleteNoteTestSuite) Test_2_Fail_NotOwner() {
	// mock data
	s.mockNoteRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Note{
			ID:     database.NewStringID(),
			UserID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	_, err := s.handler.DeleteNote(ctx, database.NewStringID(), database.NewStringID(), dto.DeleteNoteRequest{})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Common.NotFound, err)
}

//
// END OF CASES
//

func TestDeleteNoteTestSuite(t *testing.T) {
	suite.Run(t, new(deleteNoteTestSuite))
}
