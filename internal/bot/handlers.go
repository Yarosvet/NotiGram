package bot

import (
	"errors"
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

var ParseError = errors.New("failed to parse command")

func commandSubscribe(cmd string, msg *tgbotapi.Message, deps *HandlerDeps) error {
	args := stdstrings.SplitN(cmd, "-", 2)
	if len(args) < 2 {
		return ParseError
	}
	channelID := args[1]
	err := storage.Subscribe(msg.Chat.ID, channelID, deps.Logger, deps.RedisConfig)
	if err != nil {
		deps.Logger.Error("Failed to subscribe user", zap.Error(err))
		return err
	}
	_, err = deps.Bot.api.Send(
		tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(deps.Bot.strings.SubscribedFormat, channelID)),
	)
	return err
}

// TODO: 'unsubscribe' inline button under notification message
func commandUnsubscribe(cmd string, msg *tgbotapi.Message, deps *HandlerDeps) error {
	args := stdstrings.SplitN(cmd, "-", 2)
	if len(args) < 2 {
		return ParseError
	}
	channelID := args[1]
	err := storage.Unsubscribe(msg.Chat.ID, channelID, deps.Logger, deps.RedisConfig)
	if err != nil {
		return err
	}
	_, err = deps.Bot.api.Send(
		tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(deps.Bot.strings.SubscribedFormat, channelID)),
	)
	return err
}

func handleStart(msg *tgbotapi.Message, deps *HandlerDeps) error {
	words := stdstrings.Fields(msg.Text)
	if len(words) == 2 {
		cmd := stdstrings.Split(words[1], "-")
		switch cmd[0] {
		case "sub":
			err := commandSubscribe(words[1], msg, deps)
			if !errors.Is(err, ParseError) { // All cases when we didn't get ParseError
				return err
			}
		case "unsub":
			err := commandUnsubscribe(words[1], msg, deps)
			if !errors.Is(err, ParseError) { // All cases when we didn't get ParseError
				return err
			}
		}
	}
	resp := tgbotapi.NewMessage(msg.Chat.ID, deps.Bot.strings.WelcomeMessage)
	_, err := deps.Bot.api.Send(resp)
	return err
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
