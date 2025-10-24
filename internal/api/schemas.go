package api

type QueueMessageRequest struct {
	Message string `json:"message" binding:"required"`
}
