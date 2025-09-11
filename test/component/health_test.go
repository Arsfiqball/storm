package component

import (
	"app/test/testhelper"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/test-go/testify/suite"
)

type HealthTestSuite struct {
	suite.Suite
	testhelper.AppSuite
	handler http.HandlerFunc
}

func (s *HealthTestSuite) SetupTest() {
	s.Start(s.T())
	s.handler = testhelper.FiberToHandlerFunc(s.GetApp().Fiber.App())
}

func (s *HealthTestSuite) TearDownTest() {
	s.Stop(s.T())
}

func (s *HealthTestSuite) TestReadiness() {
	apitest.New().
		HandlerFunc(s.handler).
		Get("/readiness").
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}
