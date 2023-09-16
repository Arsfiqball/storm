package provider

import (
	"app/pkg/example"
	"context"
	"errors"
	"time"

	"github.com/Arsfiqball/talker/excode"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
)

func fiberHandleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(err)
}

func fiberReadiness(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

type FiberDeps struct {
	Example *example.Example
}

type Fiber struct {
	app *fiber.App
}

func NewFiber(fs *FiberDeps) *Fiber {
	app := fiber.New(fiber.Config{ErrorHandler: fiberHandleError})

	app.Use(recover.New())

	app.Get("/readiness", fiberReadiness)

	fs.Example.FiberRoute(app.Group("/example"))

	return &Fiber{app: app}
}

func (flc *Fiber) Serve(ctx context.Context) error {
	return flc.app.Listen(":3000")
}

func (flc *Fiber) Clean(ctx context.Context) error {
	return flc.app.Shutdown()
}

func (flc *Fiber) Readiness(ctx context.Context) error {
	return excode.HttpGetCheck("http://localhost:3000/readiness", 1*time.Second)(ctx)
}

var ProvideFiber = wire.NewSet(
	NewFiber,
	wire.Struct(new(FiberDeps), "*"),
)
