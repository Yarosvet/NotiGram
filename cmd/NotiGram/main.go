package NotiGram

import (
	"log"

	"github.com/Yarosvet/NotiGram/internal/bot"
	"github.com/Yarosvet/NotiGram/internal/config"
	"github.com/Yarosvet/NotiGram/internal/logger"
)

func Main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	err = logger.InitLogger(cfg.LogLevel, cfg.Dev)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	b, err := bot.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Error creating the bot: %v", err)
	}
	bot.Run(b, logger.Logger())
}
