package provider

import (
	"app/pkg/example"
	"context"
	"errors"
	"time"

	"github.com/Arsfiqball/talker/exco"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Fiber interface {
	App() *fiber.App
	Serve(ctx context.Context) error
	Clean(ctx context.Context) error
	Readiness(ctx context.Context) error
}

type fiberState struct {
	app *fiber.App
}

type FiberFeatureSet struct {
	Example *example.Example
}

func MakeFiber(fs FiberFeatureSet) (Fiber, error) {
	handleError := func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return ctx.Status(code).JSON(err)
	}

	handleReadiness := func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	}

	handleRecover := func(c *fiber.Ctx) error {
		return c.Next()
	}

	app := fiber.New(fiber.Config{ErrorHandler: handleError})

	app.Use(handleRecover)
	app.Get("/readiness", handleReadiness)
	app.Mount("/example", fs.Example.Fiber())

	return &fiberState{app: app}, nil
}

func (f *fiberState) App() *fiber.App {
	return f.app
}

func (f *fiberState) Serve(ctx context.Context) error {
	return f.app.Listen(":3000")
}

func (f *fiberState) Clean(ctx context.Context) error {
	return f.app.Shutdown()
}

func (f *fiberState) Readiness(ctx context.Context) error {
	return exco.HttpGetCheck("http://localhost:3000/readiness", 1*time.Second)(ctx)
}

var ProvideFiber = wire.NewSet(
	MakeFiber,
	wire.Struct(new(FiberFeatureSet), "*"),
)
