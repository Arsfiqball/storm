// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package example

import (
	"context"
)

// Injectors from wire.go:

func New(ctx context.Context, cfg Config) (*Example, error) {
	tracer := cfg.Tracer
	example := &Example{
		tracer: tracer,
	}
	return example, nil
}
