package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/vladesco/e-commerce/internal/logger"
)

type Config struct {
	LogLevel logger.LogLevel `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func LoadConfig() (config Config, err error) {

	if err = godotenv.Load(); err != nil {
		return
	}

	err = envconfig.Process("", &config)

	return
}
