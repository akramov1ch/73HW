package handler

import (
	"73HW/producer/model"
	"73HW/producer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	RabbitMQService service.RabbitMQService
}

func NewHandler(rabbitMQService service.RabbitMQService) *Handler {
	return &Handler{RabbitMQService: rabbitMQService}
}

func (h *Handler) CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RabbitMQService.PublishMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message published"})
}

func (h *Handler) GetMessages(c *gin.Context) {
	messages := []model.Message{}
	messages = append(messages, model.Message{RoutingKey: "key1", Content: "message1"})
	messages = append(messages, model.Message{RoutingKey: "key2", Content: "message2"})
	messages = append(messages, model.Message{RoutingKey: "key3", Content: "message3"})
	messages = append(messages, model.Message{RoutingKey: "key4", Content: "message4"})
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
