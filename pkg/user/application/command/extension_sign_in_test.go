package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	mockuser "github.com/namhq1989/tapnchill-server/internal/mock/user"
	"github.com/namhq1989/tapnchill-server/pkg/user/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type extensionSignInTestSuite struct {
	suite.Suite
	handler             command.ExtensionSignInHandler
	mockCtrl            *gomock.Controller
	mockUserRepository  *mockuser.MockUserRepository
	mockJwtRepository   *mockuser.MockJwtRepository
	mockQueueRepository *mockuser.MockQueueRepository
}

func (s *extensionSignInTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *extensionSignInTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)
	s.mockJwtRepository = mockuser.NewMockJwtRepository(s.mockCtrl)
	s.mockQueueRepository = mockuser.NewMockQueueRepository(s.mockCtrl)

	s.handler = command.NewExtensionSignInHandler(s.mockUserRepository, s.mockJwtRepository, s.mockQueueRepository)
}

func (s *extensionSignInTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *extensionSignInTestSuite) Test_1_Success() {
	// mock data
	var token = "access-token"

	s.mockUserRepository.EXPECT().
		ValidateAnonymousChecksum(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(true)

	s.mockUserRepository.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	s.mockJwtRepository.EXPECT().
		GenerateAccessToken(gomock.Any(), gomock.Any()).
		Return(token, nil)

	s.mockQueueRepository.EXPECT().
		CreateUserDefaultGoal(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ExtensionSignIn(ctx, dto.ExtensionSignInRequest{
		ClientID: "client-id",
		Checksum: "checksum",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), token, resp.AccessToken)
}

func (s *extensionSignInTestSuite) Test_2_Fail_InvalidChecksum() {
	// call
	s.mockUserRepository.EXPECT().
		ValidateAnonymousChecksum(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(false)

	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ExtensionSignIn(ctx, dto.ExtensionSignInRequest{
		ClientID: "",
		Checksum: "",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.BadRequest, err)
}

func (s *extensionSignInTestSuite) Test_2_Fail_InvalidClientID() {
	// call
	s.mockUserRepository.EXPECT().
		ValidateAnonymousChecksum(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(true)

	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ExtensionSignIn(ctx, dto.ExtensionSignInRequest{
		ClientID: "",
		Checksum: "checksum",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidClientID, err)
}

//
// END OF CASES
//

func TestExtensionSignInTestSuite(t *testing.T) {
	suite.Run(t, new(extensionSignInTestSuite))
}
