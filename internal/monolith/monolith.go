package monolith

import (
	"context"

	"github.com/vladesco/e-commerce/internal/config"
	"github.com/vladesco/e-commerce/internal/logger"
)

type Monolith struct {
	modules []Module
}

type ModuleConfig struct {
	config.AppConfig
	Logger *logger.Logger
}

type Module interface {
	Startup(context.Context, ModuleConfig) error
}

func (monolith *Monolith) AddModule(module *Module) *Monolith {
	monolith.modules = append(monolith.modules, *module)

	return monolith
}

func (monolith *Monolith) Bootstrap() (err error) {
	appConfig, err := config.LoadAppConfig()

	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	logger := logger.New(logger.LogConfig{Level: appConfig.LogLevel})

	for _, module := range monolith.modules {
		moduleConfig := ModuleConfig{
			appConfig,
			logger,
		}
		err = module.Startup(ctx, moduleConfig)

		if err != nil {
			cancel()

			return
		}
	}

	return
}
