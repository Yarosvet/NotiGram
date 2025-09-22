package bot

import (
	"github.com/Yarosvet/NotiGram/internal/strings"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	api     *tgbotapi.BotAPI
	strings *strings.Strings
}

func NewBot(token string, strings strings.Strings) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{api: api, strings: &strings}, nil
}

func configureBot(bot *Bot) error {
	commands := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: "/start", Description: bot.strings.StartCommandDescription},
	)
	_, err := bot.api.Request(commands)
	if err != nil {
		return err
	}
	return nil
}

func Run(b *Bot, l *zap.Logger) {
	l.Info("Starting bot", zap.String("username", b.api.Self.UserName))
	err := configureBot(b)
	if err != nil {
		l.Fatal("Failed to configure bot", zap.Error(err))
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		err := HandleUpdate(b.api, update)
		if err != nil {
			l.Error("Failed to handle update", zap.Error(err))
		}
	}

}
