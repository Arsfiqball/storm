package action

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type ChangePassword struct {
	Tracer trace.Tracer
}

func (a *ChangePassword) Do(c *fiber.Ctx) error {
	_, span := a.Tracer.Start(c.UserContext(), "user/internal/action.ChangePassword.Do")
	defer span.End()

	return c.SendString("ChangePassword")
}
