package bot

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Yarosvet/NotiGram/internal/storage"
	"go.uber.org/zap"
)

func ConsumeMessages(bot *Bot, logger *zap.Logger, redisConfig *storage.RedisConfig) {
	redis, err := storage.NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		logger.Error("Error connecting redis for consuming messages", zap.Error(err))
	}
	defer redis.Close()

	for {
		res, err := redis.BRPop(context.TODO(), time.Second, storage.MessagesListID).Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				logger.Error("Error consuming message", zap.Error(err))
			}
			continue
		}

		for _, msgStr := range res[1:] {
			var msg storage.QueuedMessage
			err := json.Unmarshal([]byte(msgStr), &msg)
			if err != nil {
				logger.Error("Error unmarshalling queued message", zap.Error(err))
				continue
			}
			print(*msg.Message)
		}
	}
}
