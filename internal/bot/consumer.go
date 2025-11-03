package bot

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Yarosvet/NotiGram/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// TODO: Add message rate limiting

func ConsumeMessages(bot *Bot, logger *zap.Logger, redisConfig *storage.RedisConfig) {
	redis, err := storage.NewRedisClient(context.TODO(), redisConfig, logger)
	if err != nil {
		logger.Error("Error connecting redis for consuming messages", zap.Error(err))
	}
	defer func() {
		err := redis.Close()
		if err != nil {
			logger.Error("Error closing redis connection", zap.Error(err))
		}
	}()

	for {
		res, err := redis.BRPop(context.TODO(), time.Second, storage.MessagesListID).Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				logger.Error("Error consuming message", zap.Error(err))
			}
			continue
		}

		for _, msgStr := range res[1:] {
			var qMsg storage.QueuedMessage
			err := json.Unmarshal([]byte(msgStr), &qMsg)
			if err != nil {
				logger.Error("Error unmarshalling queued message", zap.Error(err))
				continue
			}
			msgConfig := tgbotapi.NewMessage(qMsg.ChatID, *qMsg.Message)
			if msgConfig.ParseMode != "" {
				msgConfig.ParseMode = qMsg.ParseMode
			}
			msgConfig.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(bot.strings.UnsubscribeButton, "unsub-"+qMsg.ChannelID),
				),
			)
			_, err = bot.api.Send(msgConfig)
			if err != nil {
				logger.Error("Error sending message to user", zap.Int64("chatID", qMsg.ChatID), zap.Error(err))
				continue
			}
		}
	}
}
