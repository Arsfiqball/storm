package system

import (
	"app/internal/provider"

	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideFiber,
	provider.ProvideGorm,
	provider.ProvideMongo,
	provider.ProvideViper,
	provider.ProvideExample,
	wire.Struct(new(App), "*"),
)
