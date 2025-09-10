package action

import "github.com/gofiber/fiber/v2"

type VerifySignUpByEmail struct {
	//
}

func (a *VerifySignUpByEmail) Do(c *fiber.Ctx) error {
	return c.SendString("VerifySignUpByEmail")
}
