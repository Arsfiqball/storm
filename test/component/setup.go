package component

import (
	"app/internal"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type TestSuite struct {
	App         *internal.App
	HandlerFunc http.HandlerFunc
}

type TestCase struct {
	//
}

func SetupSuite(t *testing.T) *TestSuite {
	app, err := internal.New()
	if err != nil {
		t.Error(err)
	}

	return &TestSuite{
		App:         app,
		HandlerFunc: FiberToHandlerFunc(app.FiberApp),
	}
}

func (instance *TestSuite) TeardownSuite(t *testing.T) {
	// Drop data & shut down all test suite tools here e.g: stop the server
}

func (instance *TestSuite) SetupCase(t *testing.T) *TestCase {
	return &TestCase{
		//
	}
}

func (instance *TestCase) TeardownCase(t *testing.T) {
	// Reset test condition here e.g: remove all records
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
