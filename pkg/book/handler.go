package book

import (
	"app/pkg/stdapi"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	svc *BookService
}

func NewBookHandler(svc *BookService) *BookHandler {
	return &BookHandler{
		svc: svc,
	}
}

func (handler *BookHandler) CreateOne(c *fiber.Ctx) error {
	ctx := c.UserContext()

	payload := new(PayloadBook)
	if err := c.BodyParser(payload); err != nil {

		fmt.Println(payload)
		res := stdapi.NewResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(res.ToFiberMap())
	}

	result, err := handler.svc.CreateOne(ctx, *payload)
	if err != nil {
		return err
	}

	res := stdapi.NewResponse(200, "Success create one book", result)
	return c.Status(fiber.StatusOK).JSON(res.ToFiberMap())
}
