package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMessagesParam struct {
	From int `form:"from" binding:"required"`
	To   int `form:"to" binding:"required"`
}

type SendMessageBody struct {
	Id      int    `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
	From    int    `json:"from" binding:"required"`
	To      int    `json:"to" binding:"required"`
}

func RegisterMessageRouter(router *gin.RouterGroup) {
	service := GetMessageService()

	router.GET("/", func(c *gin.Context) {
		var param GetMessagesParam
		if err := c.BindQuery(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		messages := service.GetMessages(param.From, param.To)
		c.JSON(200, messages)
	})

	router.POST("/", func(c *gin.Context) {
		var request SendMessageBody
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		service.SendMessage(
			&SendMessageData{
				Id:      request.Id,
				Message: request.Message,
				From:    request.From,
				To:      request.To,
			},
		)
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}
