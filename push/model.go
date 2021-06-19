package push

import (
	"time"

	"github.com/zoripong/simple-chat-service/rpc"
)

type PushMessage struct {
	Content    string
	SenderId   int64
	ReceiverId int64
	SendAt     time.Time
}

func (message *PushMessage) ToPushInstantlyReuest() *rpc.PushInstantlyRequest {
	return &rpc.PushInstantlyRequest{
		Content:    message.Content,
		SenderId:   message.SenderId,
		ReceiverId: message.ReceiverId,
		SendAt:     message.SendAt.Unix(),
	}
}
