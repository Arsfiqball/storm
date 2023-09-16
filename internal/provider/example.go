package provider

import (
	"app/pkg/example"
	"context"
)

func ProvideExample(ctx context.Context) (*example.Example, error) {
	return example.New(ctx, example.Config{})
}
