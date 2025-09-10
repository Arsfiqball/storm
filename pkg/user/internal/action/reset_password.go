package action

import "github.com/gofiber/fiber/v2"

type ResetPassword struct {
	//
}

func (a *ResetPassword) Do(c *fiber.Ctx) error {
	return c.SendString("ResetPassword")
}
