package message

import (
	"github.com/gin-gonic/gin"
)

func RegisterMessageRouter(router *gin.RouterGroup) {
	router.POST("/", SendMessage)
}