package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	mockcommon "github.com/namhq1989/tapnchill-server/internal/mock/common"
	"github.com/namhq1989/tapnchill-server/pkg/common/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getQuoteTestSuite struct {
	suite.Suite
	handler               query.GetQuoteHandler
	mockCtrl              *gomock.Controller
	mockQuoteRepository   *mockcommon.MockQuoteRepository
	mockCachingRepository *mockcommon.MockCachingRepository
}

func (s *getQuoteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getQuoteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQuoteRepository = mockcommon.NewMockQuoteRepository(s.mockCtrl)
	s.mockCachingRepository = mockcommon.NewMockCachingRepository(s.mockCtrl)

	s.handler = query.NewGetQuoteHandler(s.mockQuoteRepository, s.mockCachingRepository)
}

func (s *getQuoteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getQuoteTestSuite) Test_1_Success() {
	// mock data
	s.mockCachingRepository.EXPECT().
		GetLatestQuote(gomock.Any()).
		Return(nil, nil)

	s.mockQuoteRepository.EXPECT().
		FindLatest(gomock.Any()).
		Return(&domain.Quote{
			ID: database.NewStringID(),
		}, nil)

	s.mockCachingRepository.EXPECT().
		SetLatestQuote(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetQuote(ctx, database.NewStringID(), dto.GetQuoteRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *getQuoteTestSuite) Test_1_Success_FoundInCaching() {
	// mock data
	s.mockCachingRepository.EXPECT().
		GetLatestQuote(gomock.Any()).
		Return(&domain.Quote{
			ID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetQuote(ctx, database.NewStringID(), dto.GetQuoteRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

//
// END OF CASES
//

func TestGetQuoteTestSuite(t *testing.T) {
	suite.Run(t, new(getQuoteTestSuite))
}
