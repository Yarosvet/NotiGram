package NotiGram

import (
	"log"

	"github.com/Yarosvet/NotiGram/internal/api"
	"github.com/Yarosvet/NotiGram/internal/bot"
	"github.com/Yarosvet/NotiGram/internal/config"
	"github.com/Yarosvet/NotiGram/internal/logger"
	"github.com/Yarosvet/NotiGram/internal/storage"
	"github.com/Yarosvet/NotiGram/internal/strings"
	"github.com/gin-gonic/gin"
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

	strs, err := strings.Load(cfg.StringsConfig)
	if err != nil {
		log.Fatalf("Error loading strings: %v", err)
	}

	b, err := bot.NewBot(cfg.TelegramToken, *strs)
	if err != nil {
		log.Fatalf("Error creating the bot: %v", err)
	}

	redisConfig := storage.DefaultRedisConfig(cfg.RedisUrl)
	go bot.Run(b, logger.Logger(), &redisConfig)

	if cfg.Dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	api.Run(logger.Logger(), &redisConfig)
}
