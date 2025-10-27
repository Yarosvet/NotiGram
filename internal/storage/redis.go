package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisConfig struct {
	Url         string
	MaxRetries  int
	Timeout     uint // in milliseconds
	DialTimeout uint // in milliseconds
}

func NewRedisClient(ctx context.Context, cfg *RedisConfig, logger *zap.Logger) (*redis.Client, error) {
	options, err := redis.ParseURL(cfg.Url)
	if err != nil {
		return nil, err
	}
	options.MaxRetries = cfg.MaxRetries
	options.ReadTimeout = time.Duration(cfg.Timeout) * time.Millisecond
	options.WriteTimeout = time.Duration(cfg.Timeout) * time.Millisecond
	options.DialTimeout = time.Duration(cfg.DialTimeout) * time.Millisecond
	db := redis.NewClient(options)
	if err := db.Ping(ctx).Err(); err != nil {
		logger.Fatal("failed to connect to redis server: %s\n", zap.Error(err))
		return nil, err
	}
	return db, nil
}

func DefaultRedisConfig(url string) RedisConfig {
	return RedisConfig{
		Url:         url,
		MaxRetries:  3,
		Timeout:     3000,
		DialTimeout: 5000,
	}
}
