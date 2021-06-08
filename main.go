package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongo/message"
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