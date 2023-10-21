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
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)

type App struct {
	fiber *provider.Fiber
}

func (a *App) Start(ctx context.Context) error {
	exec := exco.Sequential(
		exco.Parallel(
			a.fiber.Serve,
		),
	)

	return exec(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	exec := exco.Sequential(
		a.fiber.Clean,
	)

	return exec(ctx)
}

func (a *App) Live(ctx context.Context) error {
	return nil
}

func (a *App) Ready(ctx context.Context) error {
	exec := exco.Parallel(
		a.fiber.Readiness,
	)

	return exec(ctx)
}
