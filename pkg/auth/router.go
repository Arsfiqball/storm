package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoute(h IHandler, router fiber.Router) {
	router.Post("/login", h.Login)
}
