package store

import (
	"context"

	"github.com/vladesco/e-commerce/internal/config"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/internal/logger"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/repository"
	"github.com/vladesco/e-commerce/store/internal/server"
)

type StoreModule struct{}

func (module *StoreModule) Bootstrap(ctx context.Context, logger *logger.Logger, config config.ModuleConfig) {

	logger.Info().Msgf("listening port %d", config.Port)

	storeRepository := repository.NewStoreRepository(logger)
	productRepository := repository.NewProductRepository(logger)
	domainDispatcher := ddd.NewEventDispatcher()

	application := application.NewApplication(storeRepository, productRepository, domainDispatcher)

	server.RegisterServer(logger, application, &server.ServerConfig{Port: config.Port})
}
