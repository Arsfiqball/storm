package resource

import "github.com/gofiber/fiber/v2"

type One struct {
	//
}

func (r *One) Get(c *fiber.Ctx) error {
	return c.SendString("UserOne")
}

func (r *One) Patch(c *fiber.Ctx) error {
	return c.SendString("UserOnePatch")
}

func (r *One) Delete(c *fiber.Ctx) error {
	return c.SendString("UserOneDelete")
}
