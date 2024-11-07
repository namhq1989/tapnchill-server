package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockcommon "github.com/namhq1989/tapnchill-server/internal/mock/common"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/worker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getRandomQuoteTestSuite struct {
	suite.Suite
	handler                   worker.GetRandomQuoteHandler
	mockCtrl                  *gomock.Controller
	mockQuoteRepository       *mockcommon.MockQuoteRepository
	mockExternalApiRepository *mockcommon.MockExternalApiRepository
}

func (s *getRandomQuoteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getRandomQuoteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQuoteRepository = mockcommon.NewMockQuoteRepository(s.mockCtrl)
	s.mockExternalApiRepository = mockcommon.NewMockExternalApiRepository(s.mockCtrl)

	s.handler = worker.NewGetRandomQuoteHandler(s.mockQuoteRepository, s.mockExternalApiRepository)
}

func (s *getRandomQuoteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getRandomQuoteTestSuite) Test_1_Success() {
	// mock data
	s.mockExternalApiRepository.EXPECT().
		GetRandomQuote(gomock.Any()).
		Return(&domain.Quote{
			ID: database.NewStringID(),
		}, nil)

	s.mockQuoteRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.GetRandomQuote(ctx, domain.QueueGetRandomQuotePayload{})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestGetRandomQuoteTestSuite(t *testing.T) {
	suite.Run(t, new(getRandomQuoteTestSuite))
}
