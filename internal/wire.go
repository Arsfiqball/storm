//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
)

func New() (*App, error) {
	wire.Build(AppSet)
	return &App{}, nil
}
