package bot

import (
	"fmt"
	stdstrings "strings"

	"github.com/Yarosvet/NotiGram/internal/storage"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type HandlerDeps struct {
	Bot         *Bot
	Logger      *zap.Logger
	RedisConfig *storage.RedisConfig
}

func handleStart(msg *tgbotapi.Message, deps *HandlerDeps) error {
	args := stdstrings.Fields(msg.Text)
	if len(args) > 1 {
		channelID := args[1]
		err := storage.Subscribe(msg.Chat.ID, channelID, deps.Logger, deps.RedisConfig)
		if err != nil {
			return err
		}
		_, err = deps.Bot.api.Send(
			tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(deps.Bot.strings.SubscribedFormat, channelID)),
		)
		if err != nil {
			return err
		}
	} else {
		resp := tgbotapi.NewMessage(msg.Chat.ID, deps.Bot.strings.WelcomeMessage)
		_, err := deps.Bot.api.Send(resp)
		return err
	}
	return nil
}

func HandleUpdate(update tgbotapi.Update, deps *HandlerDeps) error {
	deps.Logger.Debug("Received update", zap.Any("update", update))
	if update.Message != nil && update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			return handleStart(update.Message, deps)
		}
	}
	return nil
}
