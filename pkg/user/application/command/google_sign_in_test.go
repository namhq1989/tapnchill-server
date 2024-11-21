package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/domain"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type googleSignInTestSuite struct {
	suite.Suite
	handler            command.GoogleSignInHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockuser.MockUserRepository
	mockJwtRepository  *mockuser.MockJwtRepository
	mockSSORepository  *mockuser.MockSSORepository
}

func (s *googleSignInTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *googleSignInTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockJwtRepository = mockuser.NewMockJwtRepository(s.mockCtrl)
	s.mockSSORepository = mockuser.NewMockSSORepository(s.mockCtrl)

	s.handler = command.NewGoogleSignInHandler(s.mockUserRepository, s.mockSSORepository, s.mockJwtRepository)
}

func (s *googleSignInTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *googleSignInTestSuite) Test_1_Success() {
	// mock data
	var token = "access-token"

	s.mockSSORepository.EXPECT().
		VerifyGoogleToken(gomock.Any(), gomock.Any()).
		Return(&domain.SSOGoogleUser{
			UID:   "id",
			Name:  "name",
			Email: "email",
		}, nil)

	s.mockUserRepository.EXPECT().
		FindByEmail(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockJwtRepository.EXPECT().
		GenerateAccessToken(gomock.Any(), gomock.Any()).
		Return(token, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GoogleSignIn(ctx, dto.GoogleSignInRequest{
		Token: "google token",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), token, resp.AccessToken)
}

//
// END OF CASES
//

func TestGoogleSignInTestSuite(t *testing.T) {
	suite.Run(t, new(googleSignInTestSuite))
}
