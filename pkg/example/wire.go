//go:build wireinject
// +build wireinject

package example

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context, cfg Config) (*Example, error) {
	wire.Build(RegisterSet)
	return &Example{}, nil
}
