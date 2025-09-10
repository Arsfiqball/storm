package action

import "github.com/gofiber/fiber/v2"

type SignOut struct {
	//
}

func (a *SignOut) Do(c *fiber.Ctx) error {
	return c.SendString("SignOut")
}
