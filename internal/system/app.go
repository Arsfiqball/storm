package system

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	FiberApp    *fiber.App
	GormDb      *gorm.DB
	ViperConfig *viper.Viper
}

func (a *App) Serve(ctx context.Context) error {
	return a.FiberApp.Listen(a.ViperConfig.GetString("address"))
}

func (a *App) Clean(ctx context.Context) error {
	a.FiberApp.Shutdown()

	sqlDB, err := a.GormDb.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
