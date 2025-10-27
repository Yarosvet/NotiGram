package api

import (
	"github.com/Yarosvet/NotiGram/internal/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func configureRoutes(e *gin.Engine) {
	e.POST("/queue/:channelID", queueMessage)
}

func configureMiddleware(e *gin.Engine, logger *zap.Logger, redisConfig *storage.RedisConfig) {
	e.Use(func(c *gin.Context) {
		c.Set("logger", logger)
		c.Set("redisConfig", redisConfig)
		c.Next()
	})
}

func Run(logger *zap.Logger, redisConfig *storage.RedisConfig) {
	r := gin.Default()
	configureMiddleware(r, logger, redisConfig)
	configureRoutes(r)
	err := r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080  // TODO: make address configurable
	if err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
