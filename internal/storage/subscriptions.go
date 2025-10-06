package storage

import (
	"context"
	"strconv"

	"go.uber.org/zap"
)

func Subscribe(chatID int64, channelID string, logger *zap.Logger, redisConfig *RedisConfig) error {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return err
	}
	redis.Set(context.TODO(), "sub:"+strconv.FormatInt(chatID, 10)+":"+channelID, 1, 0)
	logger.Info("User subscribed", zap.Int64("userID", chatID), zap.String("channelID", channelID))
	return nil
}

func List(chatID int64, logger *zap.Logger, redisConfig *RedisConfig) (*[]string, error) {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return nil, err
	}
	keys, err := redis.Keys(context.TODO(), "sub:"+strconv.FormatInt(chatID, 10)+":*").Result()
	if err != nil {
		return nil, err
	}
	channels := make([]string, len(keys))
	for i, key := range keys {
		channels[i] = key[len("sub:"+strconv.FormatInt(chatID, 10)+":"):]
	}
	logger.Info("User channel list retrieved", zap.Int64("userID", chatID), zap.Strings("channels", channels))
	return &channels, nil
}
