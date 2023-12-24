package provider

import (
	"app/pkg/example"
	"context"
)

// ProvideExample is a Wire provider function that returns a *example.Example.
func ProvideExample(ctx context.Context) (*example.Example, error) {
	return example.New(ctx, example.Config{})
}
