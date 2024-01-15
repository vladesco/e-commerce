package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/vladesco/e-commerce/internal/logger"
)

type AppConfig struct {
	LogLevel logger.LogLevel `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func LoadAppConfig() (appConfig AppConfig, err error) {

	if err = godotenv.Load(); err != nil {
		return
	}

	err = envconfig.Process("", &appConfig)

	return
}
