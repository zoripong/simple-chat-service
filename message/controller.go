package message

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageBody struct {
	Id	int `json:"id" binding:"required"`
	Message	string `json:"message" binding:"required"`
	From int `json:"from" binding:"required"`
	To int `json:"to" binding:"required"`
}

func SendMessage(c *gin.Context) {
	var request SendMessageBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// FIXME service 거쳐서 저장할 수 있게끔 수정해야 함
	repository := GetMessageRespository()
	repository.Save(
		&Message{
			Id:         request.Id,
			Message:    request.Message,
			From:       request.From,
			To:         request.To,
			SendAt:     time.Now(),
			ReceivedAt: time.Now(),
		},
	)
	c.JSON(200, gin.H{
		"sentAt": time.Now(),
	})
}
