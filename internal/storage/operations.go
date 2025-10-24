package storage

import (
	"context"
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type QueuedMessage struct {
	ChatID  int64   `json:"chat_id"`
	Message *string `json:"message"`
}

const MessagesListID = "messages"

func Subscribe(chatID int64, channelID string, logger *zap.Logger, redisConfig *RedisConfig) error {
	redis, err := NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		return err
	}
	err = redis.SAdd(context.TODO(), "sub:"+channelID, chatID).Err()
	if err != nil {
		return err
	}
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
	members, err := redis.SMembers(context.TODO(), "sub:"+channelID).Result()
	if err != nil {
		return err
	}
	for _, chatID := range members {
		intChatID, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			logger.Error("Failed to parse chatID from redis", zap.String("chatID", chatID), zap.Error(err))
			continue
		}
		qMsg := QueuedMessage{ChatID: intChatID, Message: msg}
		jsonMsg, err := json.Marshal(qMsg)
		if err != nil {
			logger.Error("Failed to marshal queued message", zap.Int64("chatID", intChatID), zap.Error(err))
			continue
		}
		err = redis.LPush(context.TODO(), MessagesListID, jsonMsg).Err()
		if err != nil {
			logger.Error("Failed to queue message", zap.Int64("chatID", intChatID), zap.Error(err))
			continue
		}
	}
	return redis.Close()
}
