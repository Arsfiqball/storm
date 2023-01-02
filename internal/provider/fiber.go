package provider

import (
	"app/pkg/book"
	"app/pkg/stdapi"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
)

type FiberHandlerSet struct {
	BookHandler *book.BookHandler
}

func NewFiber(hs *FiberHandlerSet) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			res := stdapi.NewResponse(code, err.Error(), nil)
			return ctx.Status(code).JSON(res.ToFiberMap())
		},
	})

	app.Use(recover.New())

	app.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Register all router
	book.Route(app.Group("/v1/book"), hs.BookHandler)

	return app
}

var ProvideFiber = wire.NewSet(
	NewFiber,
	wire.Struct(new(FiberHandlerSet), "*"),
)
