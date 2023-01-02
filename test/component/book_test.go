package component

import (
	"app/test/testhelper"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
)

type BookTestSuite struct {
	suite.Suite
	testhelper.AppSuite
	handler http.HandlerFunc
}

func (s *BookTestSuite) SetupTest() {
	s.Start(s.T())
	s.ExecSQL(s.T(), "DELETE FROM books")
	s.handler = testhelper.FiberToHandlerFunc(s.GetApp().FiberApp)
}

func (s *BookTestSuite) TearDownTest() {
	s.Stop(s.T())
}

func (s *BookTestSuite) TestReadiness() {
	type scenarioTemplate struct {
		name        string
		payload     map[string]interface{}
		checkStatus int
		checkJson   func(*http.Response, *http.Request) error
	}

	scenarios := []scenarioTemplate{
		{
			name:        "Insert a book",
			payload:     map[string]interface{}{},
			checkStatus: 200,
			checkJson: jsonpath.Chain().
				End(),
		},
	}

	for _, scenario := range scenarios {
		s.Run(scenario.name, func() {
			var body interface{}

			bodyBytes, err := json.Marshal(scenario.payload)
			if err != nil {
				s.Fail(err.Error())
			}

			apitest.New().
				HandlerFunc(s.handler).
				Post("/book/one").
				Body(string(bodyBytes)).
				Expect(s.T()).
				Status(scenario.checkStatus).
				Assert(scenario.checkJson).
				End().
				JSON(&body)

			// DEBUG by adding verbose flag (-v)
			testhelper.JsonPrint(body)
		})
	}
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
