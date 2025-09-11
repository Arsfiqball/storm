package system

import (
	"app/internal/provider"

	"github.com/Arsfiqball/codec/talker"
	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideOtel,
	provider.ProvideSlog,
	provider.ProvideFiber,
	provider.ProvideGorm,
	provider.ProvideWatermill,
	provider.ProvideWork,
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)

type App struct {
	Otel      provider.Otel
	Fiber     provider.Fiber
	Gorm      provider.Gorm
	Watermill provider.Watermill
	Work      provider.Work
	Slog      provider.Slog
}

func (a *App) Serve() talker.Process {
	return talker.Process{
		Start: talker.Parallel(
			a.Fiber.Serve,
			a.Watermill.Serve,
			a.Work.Start,
		),
		Ready: talker.Parallel(
			a.Fiber.Readiness,
			a.Gorm.Ping,
		),
		Stop: talker.Sequential(
			talker.Parallel(
				a.Fiber.Clean,
				a.Watermill.Clean,
				a.Work.Stop,
			),
			a.Gorm.Close,
		),
	}
}
