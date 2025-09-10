package action

import "github.com/gofiber/fiber/v2"

type ForgotPassword struct {
	//
}

func (a *ForgotPassword) Do(c *fiber.Ctx) error {
	return c.SendString("ForgotPassword")
}
