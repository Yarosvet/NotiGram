package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{api: api}, nil
}

func Run(b *Bot, l *zap.Logger) {
	l.Info("Bot started", zap.String("username", b.api.Self.UserName))
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
