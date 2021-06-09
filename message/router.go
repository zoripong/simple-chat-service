package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMessagesParam struct {
	From int `form:"from" binding:"required"`
	To int `form:"to" binding:"required"`
}

type SendMessageBody struct {
	Id      int    `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
	From    int    `json:"from" binding:"required"`
	To      int    `json:"to" binding:"required"`
}

func RegisterMessageRouter(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		var param GetMessagesParam
		if err := c.BindQuery(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, "Message sent to %d from %d", param.To, param.From)
	})

	router.POST("/", func(c *gin.Context) {
		var request SendMessageBody
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		service := GetMessageService()
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
