package main

import (
	"context"

	"github.com/vladesco/e-commerce/internal/monolith"
)

type StoreModule struct{}

func (module *StoreModule) Startup(ctx context.Context, config monolith.ModuleConfig) error {
	return nil
}
