//go:build wireinject
// +build wireinject

package user

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context, cfg Config) (*User, error) {
	wire.Build(RegisterSet)
	return &User{}, nil
}
