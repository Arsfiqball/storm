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

func (a *App) ServeOnlyHTTP() talker.Process {
	return talker.Process{
		Start: a.Fiber.Serve,
		Ready: talker.Parallel(
			a.Fiber.Readiness,
			a.Gorm.Ping,
		),
		Stop: talker.Sequential(
			a.Fiber.Clean,
			a.Gorm.Close,
		),
	}
}

func (a *App) ServeOnlyListener() talker.Process {
	return talker.Process{
		Start: a.Watermill.Serve,
		Ready: a.Gorm.Ping,
		Stop: talker.Sequential(
			a.Watermill.Clean,
			a.Gorm.Close,
		),
	}
}

func (a *App) ServeOnlyWorker() talker.Process {
	return talker.Process{
		Start: a.Work.Start,
		Ready: a.Gorm.Ping,
		Stop: talker.Sequential(
			a.Work.Stop,
			a.Gorm.Close,
		),
	}
}
