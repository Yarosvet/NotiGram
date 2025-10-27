package api

import "fmt"

type QueueMessageRequest struct {
	Message   string `json:"message" binding:"required"`
	ParseMode string `json:"parse_mode" binding:"required"`
}

func (q *QueueMessageRequest) Validate() error {
	if q == nil {
		return fmt.Errorf("request is nil")
	}
	switch q.ParseMode {
	case "Markdown", "HTML", "MarkdownV2":
		return nil
	default:
		return fmt.Errorf("invalid parse_mode: %q (allowed: Markdown, MarkdownV2, HTML)", q.ParseMode)
	}
}
