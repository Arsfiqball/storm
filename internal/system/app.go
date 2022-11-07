package system

import (
	"app/internal/provider"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var RegisterSet = wire.NewSet(
	NewApp,
	provider.ProvideFiber,
	provider.ProvideViper,
)

type App struct {
	fiberApp    *fiber.App
	viperConfig *viper.Viper
}

func NewApp(fiberApp *fiber.App, viperConfig *viper.Viper) *App {
	return &App{
		fiberApp:    fiberApp,
		viperConfig: viperConfig,
	}
}

func (a *App) Serve(ctx context.Context) error {
	return a.fiberApp.Listen(a.viperConfig.GetString("address"))
}

func (a *App) GetFiber() *fiber.App {
	return a.fiberApp
}

func (a *App) Clean(ctx context.Context) error {
	return a.fiberApp.Shutdown()
}
