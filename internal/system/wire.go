//go:build wireinject
// +build wireinject

package system

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context) (*App, error) {
	wire.Build(RegisterSet)
	return &App{}, nil
}
