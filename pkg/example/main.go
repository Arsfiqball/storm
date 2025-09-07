package example

import (
	"errors"

	"github.com/Arsfiqball/codec/flame"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

type Config struct {
	Tracer trace.Tracer
}

type Example struct {
	tracer trace.Tracer
}

func (e *Example) Fiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		_, span := e.tracer.Start(c.UserContext(), "example.handler")
		defer span.End()

		return c.SendString("Hello, World 👋!")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		_, span := e.tracer.Start(c.UserContext(), "example.errorHandler")
		defer span.End()

		return flame.Unexpected(errors.New("an unexpected error occurred"))
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		_, span := e.tracer.Start(c.UserContext(), "example.panicHandler")
		defer span.End()

		panic("this is a panic")
	})

	return app
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Example), "*"),
	wire.FieldsOf(new(Config), "Tracer"),
)
