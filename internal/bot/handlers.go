package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	resp := tgbotapi.NewMessage(msg.Chat.ID, "Hello!")
	_, err := bot.Send(resp)
	return err
}

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	if update.Message != nil && update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			err := handleStart(bot, update.Message)
			return err
		}
	}
	return nil
}
