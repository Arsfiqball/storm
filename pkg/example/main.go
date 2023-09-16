package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Config struct{}

type Example struct{}

func (e *Example) FiberRoute(router fiber.Router) {
	//
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Example), "*"),
	// wire.FieldsOf(new(Config)),
)
