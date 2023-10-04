package example

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Config struct{}

type Example struct {
	handler *Handler
}

func (e *Example) FiberRoute(router fiber.Router) {
	router.Get("/", e.handler.GetOneUser)
}

type Handler struct {
	//
}

func (h *Handler) GetOneUser(c *fiber.Ctx) error {
	return nil
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Example), "*"),
	// wire.FieldsOf(new(Config)),
)
