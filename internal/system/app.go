package system

import (
	"app/internal/provider"
	"context"

	"github.com/Arsfiqball/talker/excode"
	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideViper,
	provider.ProvideFiber,
	provider.ProvideGorm,
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)

type App struct {
	fiber *provider.Fiber
	gorm  *provider.Gorm
}

func (a *App) Serve(ctx context.Context) error {
	exec := excode.Sequential(
		excode.Parallel(
			a.fiber.Serve,
		),
	)

	return exec(ctx)
}

func (a *App) Clean(ctx context.Context) error {
	exec := excode.Sequential(
		a.fiber.Clean,
		a.gorm.Clean,
	)

	return exec(ctx)
}

func (a *App) Liveness(ctx context.Context) error {
	return nil
}

func (a *App) Readiness(ctx context.Context) error {
	exec := excode.Parallel(
		a.fiber.Readiness,
	)

	return exec(ctx)
}
