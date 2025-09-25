package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	LogLevel      string `env:"LOG_LEVEL" def:"info"`
	RedisUrl      string `env:"REDIS_URL"`
	Dev           bool   `env:"DEV" def:"false"`
	StringsConfig string `env:"STRINGS_CONFIG" def:""`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
