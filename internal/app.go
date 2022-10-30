package internal

import (
	"app/pkg/auth"
	"app/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var AppSet = wire.NewSet(
	NewApp,
	NewGormDB,
	auth.NewHandler,
	wire.Bind(new(auth.IHandler), new(*auth.Handler)),
	auth.NewUserRepo,
	wire.Bind(new(auth.IUserRepo), new(*auth.UserRepo)),
)

type App struct {
	FiberApp *fiber.App
}

func NewApp(db *gorm.DB, authHandler auth.IHandler) *App {
	user.Migrate(db)

	fiberApp := fiber.New()

	fiberApp.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	fiberApp.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	apiV1 := fiberApp.Group("/v1")

	apiAuth := apiV1.Group("/auth")
	auth.RegisterRoute(authHandler, apiAuth)

	return &App{
		FiberApp: fiberApp,
	}
}
