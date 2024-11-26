package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mocknote "github.com/namhq1989/tapnchill-server/internal/mock/note"
	"github.com/namhq1989/tapnchill-server/pkg/note/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type syncTestSuite struct {
	suite.Suite
	handler            query.SyncHandler
	mockCtrl           *gomock.Controller
	mockNoteRepository *mocknote.MockNoteRepository
}

func (s *syncTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *syncTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockNoteRepository = mocknote.NewMockNoteRepository(s.mockCtrl)

	s.handler = query.NewSyncHandler(s.mockNoteRepository)
}

func (s *syncTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *syncTestSuite) Test_1_Success() {
	// mock data
	s.mockNoteRepository.EXPECT().
		Sync(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]domain.Note{
			{ID: database.NewStringID()},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.Sync(ctx, database.NewStringID(), dto.SyncRequest{
		LastUpdatedAt: "",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Notes))
}

//
// END OF CASES
//

func TestSyncTestSuite(t *testing.T) {
	suite.Run(t, new(syncTestSuite))
}
