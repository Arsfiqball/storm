package component_test

import (
	"app/test/component"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestHealth(t *testing.T) {
	testSuite := component.SetupSuite(t)
	defer testSuite.TeardownSuite(t)

	scenarios := []struct {
		title          string
		path           string
		expectedStatus int
	}{
		{"Readiness", "/readiness", http.StatusOK},
		{"Liveness", "/liveness", http.StatusOK},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.title, func(t *testing.T) {
			testCase := testSuite.SetupCase(t)
			defer testCase.TeardownCase(t)

			apitest.New().
				HandlerFunc(testSuite.HandlerFunc).
				Get(scenario.path).
				Expect(t).
				Status(scenario.expectedStatus).
				End()
		})
	}
}
