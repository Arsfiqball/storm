package provider

import (
	"github.com/gofiber/fiber/v2"
)

func ProvideFiber() *fiber.App {
	app := fiber.New()

	app.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	return app
}
