package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoripong/simple-chat-service/message"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	messageRouterGroup := r.Group("/message")
	message.RegisterMessageRouter(messageRouterGroup)

	r.Run()
}
