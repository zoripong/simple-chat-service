package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageBody struct {
	Id	int `json:"id" binding:"required"`
	Message	string `json:"message" binding:"required"`
	From int `json:"from" binding:"required"`
	To int `json:"to" binding:"required"`
}

func RegisterMessageRouter(router *gin.RouterGroup) {
	router.POST("/", func (c *gin.Context) {
		var request SendMessageBody
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		service := GetMessageService()
		service.SendMessage(
			&SendMessageData{
				Id: request.Id,
				Message: request.Message,
				From: request.From,
				To: request.To,
			},
		)
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}
