package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Config struct {
	//
}

type Example struct{}

func (e *Example) Fiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Example), "*"),
	// wire.FieldsOf(new(Config)),
)
