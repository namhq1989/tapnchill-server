package query_test

import (
	"testing"

	"github.com/namhq1989/tapnchill-server/pkg/user/application/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type getSubscriptionPlansTestSuite struct {
	suite.Suite
	handler query.GetSubscriptionPlansHandler
}

func (s *getSubscriptionPlansTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getSubscriptionPlansTestSuite) setupApplication() {
	s.handler = query.NewGetSubscriptionPlansHandler()
}

//
// CASES
//

func (s *getSubscriptionPlansTestSuite) Test_1_Success() {
	assert.Equal(s.T(), 1, 1)
}

//
// END OF CASES
//

func TestGetSubscriptionPlansTestSuite(t *testing.T) {
	suite.Run(t, new(getSubscriptionPlansTestSuite))
}
