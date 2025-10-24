package storage

import (
	"context"
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type QueuedMessage struct {
	ChannelID string  `json:"channel_id"`
	Message   *string `json:"message"`
}

const MessagesListID = "messages"

func Subscribe(chatID int64, channelID string, logger *zap.Logger, redisConfig *RedisConfig) error {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return err
	}
	redis.Set(context.TODO(), "sub:"+strconv.FormatInt(chatID, 10)+":"+channelID, 1, 0)
	logger.Info("User subscribed", zap.Int64("userID", chatID), zap.String("channelID", channelID))
	return redis.Close()
}

func Unsubscribe(chatID int64, channelID string, logger *zap.Logger, redisConfig *RedisConfig) error {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return err
	}
	redis.Del(context.TODO(), "sub:"+strconv.FormatInt(chatID, 10)+":"+channelID)
	logger.Info("User unsubscribed", zap.Int64("userID", chatID), zap.String("channelID", channelID))
	return redis.Close()
}

func QueueMessage(channelID string, msg *string, logger *zap.Logger, redisConfig *RedisConfig) error {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return err
	}
	qMsg := QueuedMessage{ChannelID: channelID, Message: msg}
	jsonMsg, err := json.Marshal(qMsg)
	if err != nil {
		return err
	}
	redis.LPush(context.TODO(), MessagesListID, jsonMsg)
	return redis.Close()
}
