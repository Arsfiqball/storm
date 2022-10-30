package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	Login(*fiber.Ctx) error
}

type Handler struct {
	userRepo IUserRepo
}

func NewHandler(userRepo IUserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}

type LoginRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cred := new(LoginRequest)

	if err := c.BodyParser(cred); err != nil {
		return err
	}

	user, err := h.userRepo.FindByEmail(ctx, cred.Email)
	if err != nil {
		return nil
	}

	fmt.Println(user)

	return c.JSON(&fiber.Map{
		"code":      200,
		"data":      nil,
		"messageId": "SUCCESS",
	})
}
