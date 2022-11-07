package component

import (
	"app/internal/system"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
)

type HealthTestSuite struct {
	suite.Suite
	app     *system.App
	appCtx  context.Context
	handler http.HandlerFunc
}

func (s *HealthTestSuite) SetupTest() {
	ctx := context.Background()
	app, err := system.New(ctx)
	if err != nil {
		s.Fail("Failed to start server")
	}

	s.app = app
	s.appCtx = ctx
	s.handler = FiberToHandlerFunc(s.app.GetFiber())
}

func (s *HealthTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(s.appCtx, 30*time.Second)
	defer cancel()

	err := s.app.Clean(ctx)
	if err != nil {
		s.Fail("Failed to stop server")
	}
}

func (s *HealthTestSuite) TestReadiness() {
	apitest.New().
		HandlerFunc(s.handler).
		Get("/readiness").
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}
