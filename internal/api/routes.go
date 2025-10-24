package api

import (
	"net/http"

	"github.com/Yarosvet/NotiGram/internal/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func queueMessage(c *gin.Context) {
	var input QueueMessageRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := storage.QueueMessage(
		c.Param("channelID"),
		&input.Message,
		c.MustGet("logger").(*zap.Logger),
		c.MustGet("redisConfig").(*storage.RedisConfig),
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.MustGet("logger").(*zap.Logger).Error("Failed to queue message", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{"ok": true})
}
