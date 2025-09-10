package content

import "github.com/gofiber/fiber/v2"

type CurrentSession struct {
	//
}

func (cs *CurrentSession) Fetch(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"session": fiber.Map{
			"id":     "session_123",
			"userId": "user_456",
			"status": "active",
		},
	})
}
