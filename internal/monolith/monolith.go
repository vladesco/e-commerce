package monolith

import (
	"context"
	"net"

	"github.com/vladesco/e-commerce/internal/config"
	"github.com/vladesco/e-commerce/internal/logger"
)

type Module interface {
	Bootstrap(context.Context, config.ModuleConfig)
}

type Monolith struct {
	modules []Module
}

func (monolith *Monolith) AddModule(module Module) *Monolith {
	monolith.modules = append(monolith.modules, module)

	return monolith
}

func (monolith *Monolith) Bootstrap() error {
	monolithConfig, err := config.LoadMonolithConfig()

	if err != nil {
		return err
	}

	logger := logger.New(logger.LogConfig{Level: monolithConfig.LogLevel})
	ctx, cancel := context.WithCancel(context.Background())

	for _, module := range monolith.modules {
		port, err := monolith.findAvailablePort()

		if err != nil {
			cancel()
			return err
		}

		go module.Bootstrap(ctx, config.ModuleConfig{
			Environment: monolithConfig.Environment,
			Port:        port,
			Logger:      logger,
		})

	}

	<-ctx.Done()

	return nil
}

func (monolith *Monolith) findAvailablePort() (int, error) {
	address, err := net.ResolveTCPAddr("tcp", ":0")

	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		return 0, err
	}

	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
