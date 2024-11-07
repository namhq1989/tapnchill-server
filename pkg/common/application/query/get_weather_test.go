package query_test

import (
	"context"
	"testing"

	apperrors "github.com/namhq1989/tapnchill-server/internal/error"

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

type getWeatherTestSuite struct {
	suite.Suite
	handler     query.GetWeatherHandler
	mockCtrl    *gomock.Controller
	mockService *mockcommon.MockService
}

func (s *getWeatherTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getWeatherTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mockcommon.NewMockService(s.mockCtrl)

	s.handler = query.NewGetWeatherHandler(s.mockService)
}

func (s *getWeatherTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getWeatherTestSuite) Test_1_Success() {
	// mock data
	var (
		city = "Da Nang"
		temp = 35.5
	)

	s.mockService.EXPECT().
		GetIpCity(gomock.Any(), gomock.Any()).
		Return(&city, nil)

	s.mockService.EXPECT().
		GetCityWeather(gomock.Any(), gomock.Any()).
		Return(&domain.Weather{
			Temp: temp,
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetWeather(ctx, database.NewStringID(), dto.GetWeatherRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), temp, resp.Weather.Temp)
}

func (s *getWeatherTestSuite) Test_2_Fail_InvalidCity() {
	// mock data
	s.mockService.EXPECT().
		GetIpCity(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetWeather(ctx, database.NewStringID(), dto.GetWeatherRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

func (s *getWeatherTestSuite) Test_2_Fail_InvalidWeather() {
	// mock data
	var (
		city = "Da Nang"
	)

	s.mockService.EXPECT().
		GetIpCity(gomock.Any(), gomock.Any()).
		Return(&city, nil)

	s.mockService.EXPECT().
		GetCityWeather(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetWeather(ctx, database.NewStringID(), dto.GetWeatherRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

//
// END OF CASES
//

func TestGetWeatherTestSuite(t *testing.T) {
	suite.Run(t, new(getWeatherTestSuite))
}
