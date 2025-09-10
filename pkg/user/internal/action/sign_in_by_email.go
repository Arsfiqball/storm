package action

import "github.com/gofiber/fiber/v2"

type SignInByEmail struct {
	//
}

func (a *SignInByEmail) Do(c *fiber.Ctx) error {
	return c.SendString("SignInByEmail")
}
