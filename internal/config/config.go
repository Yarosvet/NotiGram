package config

import (
	"github.com/caarlos0/env/v11"
)

// TODO: Add webhooks option

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	LogLevel      string `env:"LOG_LEVEL" def:"info"`
	RedisUrl      string `env:"REDIS_URL"`
	Dev           bool   `env:"DEV" def:"false"`
	StringsConfig string `env:"STRINGS_CONFIG" def:""`
	Address       string `env:"ADDRESS" def:"127.0.0.1:8080"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
