package auth_test

import (
	"app/test/component"
	"fmt"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestAuth_Login(t *testing.T) {
	testSuite := component.SetupSuite(t)
	defer testSuite.TeardownSuite(t)

	scenarios := []struct {
		title          string
		inputEmail     string
		inputPassword  string
		expectedStatus int
		expectedHasJWT bool
	}{
		{"Success", "admin@example.com", "pass1234", http.StatusOK, true},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.title, func(t *testing.T) {
			testCase := testSuite.SetupCase(t)
			defer testCase.TeardownCase(t)

			apitest.New().
				HandlerFunc(testSuite.HandlerFunc).
				Post("/v1/auth/login").
				JSON(fmt.Sprintf(`{
					"email": "%s",
					"password": "%s"
				}`, scenario.inputEmail, scenario.inputPassword)).
				Expect(t).
				Status(scenario.expectedStatus).
				End()
		})
	}
}
