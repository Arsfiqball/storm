package provider

import (
	"app/pkg/example"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Fiber interface {
	App() *fiber.App
	Serve(ctx context.Context) error
	Clean(ctx context.Context) error
	Readiness(ctx context.Context) error
}

type fiberState struct {
	app  *fiber.App
	addr string
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

	addr := viper.GetString("serve.addr")
	if addr == "" {
		addr = ":3000"
	}

	return &fiberState{app: app, addr: addr}, nil
}

func (f *fiberState) App() *fiber.App {
	return f.app
}

func (f *fiberState) Serve(ctx context.Context) error {
	return f.app.Listen(f.addr)
}

func (f *fiberState) Clean(ctx context.Context) error {
	return f.app.Shutdown()
}

func (f *fiberState) Readiness(ctx context.Context) error {
	return nil
}

var ProvideFiber = wire.NewSet(
	MakeFiber,
	wire.Struct(new(FiberFeatureSet), "*"),
)
