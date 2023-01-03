package book

import "github.com/gofiber/fiber/v2"

func Route(router fiber.Router, bookHandler *BookHandler) {
	router.Post("/one", bookHandler.CreateOne)
	router.Get("/one", bookHandler.GetOne)
}
