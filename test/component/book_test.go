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

func (s *BookTestSuite) TestCreateOne() {
	type scenarioTemplate struct {
		name        string
		payload     map[string]interface{}
		checkStatus int
		checkJson   func(*http.Response, *http.Request) error
	}

	scenarios := []scenarioTemplate{
		{
			name: "Insert a book",
			payload: map[string]interface{}{
				"title":       "Clean Architecture",
				"author":      "Uncle Bob",
				"series":      "Programming",
				"volume":      3,
				"fileUrl":     "https://book.com/clean.pdf",
				"coverUrl":    "https://book.com/cover_clean.jpg",
				"publishDate": "2023-01-01T23:11:57+07:00",
			},
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

			bodyString := string(bodyBytes)

			apitest.New().
				HandlerFunc(s.handler).
				Post("/v1/book/one").
				Header("Content-Type", "application/json").
				Body(bodyString).
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
