package system

import (
	"app/internal/provider"
	"context"

	"github.com/Arsfiqball/csverse/talker"
	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideOTEL,
	provider.ProvideSlog,
	provider.ProvideFiber,
	provider.ProvideGORM,
	provider.ProvideWatermill,
	provider.ProvideWork,
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)

type App struct {
	OTEL      provider.OTEL
	Fiber     provider.Fiber
	GORM      provider.GORM
	Watermill provider.Watermill
	Work      provider.Work
}

func (a *App) Start(ctx context.Context) error {
	exec := talker.Sequential(
		talker.Parallel(
			a.Fiber.Serve,
			a.Watermill.Serve,
			a.Work.Start,
		),
	)

	return exec(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	exec := talker.Sequential(
		talker.Parallel(
			a.Fiber.Clean,
			a.Watermill.Clean,
			a.Work.Stop,
		),
		a.GORM.Close,
	)

	return exec(ctx)
}

func (a *App) Live(ctx context.Context) error {
	return nil
}

func (a *App) Ready(ctx context.Context) error {
	exec := talker.Parallel(
		a.Fiber.Readiness,
		a.GORM.Ping,
	)

	return exec(ctx)
}
