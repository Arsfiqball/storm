package system

import (
	"app/internal/provider"

	"github.com/google/wire"
)

var RegisterSet = wire.NewSet(
	provider.ProvideFiber,
	provider.ProvideGorm,
	provider.ProvideViper,
	provider.ProvideBook,
	wire.Struct(new(App), "*"),
)
