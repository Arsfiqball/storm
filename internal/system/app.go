package system

import (
	"app/internal/provider"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var RegisterSet = wire.NewSet(
	NewApp,
	provider.ProvideFiber,
	provider.ProvideGorm,
	provider.ProvideViper,
)

type App struct {
	fiberApp    *fiber.App
	gormDb      *gorm.DB
	viperConfig *viper.Viper
}

func NewApp(fiberApp *fiber.App, gormDb *gorm.DB, viperConfig *viper.Viper) *App {
	return &App{
		fiberApp:    fiberApp,
		viperConfig: viperConfig,
		gormDb:      gormDb,
	}
}

func (a *App) Serve(ctx context.Context) error {
	return a.fiberApp.Listen(a.viperConfig.GetString("address"))
}

func (a *App) GetFiber() *fiber.App {
	return a.fiberApp
}

func (a *App) Clean(ctx context.Context) error {
	a.fiberApp.Shutdown()

	sqlDB, err := a.gormDb.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
