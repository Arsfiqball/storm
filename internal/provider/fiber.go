package provider

import (
	"app/pkg/example"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
)

type FiberFeatureSet struct {
	Example *example.Example
}

func NewFiber(fs *FiberFeatureSet) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: fiberHandleError})
	app.Use(recover.New())
	app.Get("/readiness", fiberReadiness)
	fs.Example.FiberRoute(app.Group("/example"))

	return app
}

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

var ProvideFiber = wire.NewSet(
	NewFiber,
	wire.Struct(new(FiberFeatureSet), "*"),
)
