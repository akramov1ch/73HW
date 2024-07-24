package router

import (
    "73HW/producer/handler"
    "73HW/producer/service"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    rabbitMQService := service.NewRabbitMQService()
    h := handler.NewHandler(rabbitMQService)

    r.POST("/messages", h.CreateMessage)
    r.GET("/messages", h.GetMessages)

    return r
}
