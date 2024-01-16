package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/vladesco/e-commerce/internal/logger"
)

type ModuleConfig struct {
	Environment string
	Port        int
	Logger      *logger.Logger
}

type MonolithConfig struct {
	Environment string          `envconfig:"ENVIRONMENT" default:"DEV"`
	LogLevel    logger.LogLevel `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func LoadMonolithConfig() (monolithConfig MonolithConfig, err error) {

	if err = godotenv.Load(); err != nil {
		return
	}

	err = envconfig.Process("", &monolithConfig)

	return
}
