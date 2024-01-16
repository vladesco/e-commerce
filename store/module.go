package store

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vladesco/e-commerce/internal/config"
)

type StoreModule struct{}

func (module *StoreModule) Bootstrap(ctx context.Context, config config.ModuleConfig) {
	config.Logger.Info().Msgf("listening port %d", config.Port)

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
