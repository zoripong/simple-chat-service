package message

import (
	"time"
	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context){
	repository := GetMessageRespository()
	// FIXME
	// request body를 통해 메시지 저장할 수 있도록 수정해야 함
	// service 거쳐서 저장할 수 있게끔 수정해야 함
	repository.Save(
		&Message{
			Id: 1,
			Message: "Hi!",
			From: 2,
			To: 3,
			SendAt: time.Now(),
			ReceivedAt: time.Now(),
		},
	)
	c.JSON(200, gin.H{ 
		"sentAt": time.Now(),
	})
}