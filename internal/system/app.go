package system

import (
	"app/internal/provider"
	"context"

	"github.com/Arsfiqball/talker/exco"
	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideSlog,
	provider.ProvideFiber,
	provider.ProvideGORM,
	provider.ProvideWatermill,
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)

type App struct {
	Fiber     provider.Fiber
	GORM      provider.GORM
	Watermill provider.Watermill
}

func (a *App) Start(ctx context.Context) error {
	exec := exco.Sequential(
		exco.Parallel(
			a.Fiber.Serve,
			a.Watermill.Serve,
		),
	)

	return exec(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	return exco.Sequential(
		exco.Parallel(
			a.Fiber.Clean,
			a.Watermill.Clean,
		),
		a.GORM.Close,
	)(ctx)
}

func (a *App) Live(ctx context.Context) error {
	return nil
}

func (a *App) Ready(ctx context.Context) error {
	exec := exco.Parallel(
		a.Fiber.Readiness,
	)

	return exec(ctx)
}
