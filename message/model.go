package message

import (
	"fmt"
	"strings"
	"time"

	"github.com/zoripong/simple-chat-service/public"
)

type Message struct {
	Id         int
	Message    string
	From       int
	To         int
	SendAt     time.Time
	ReceivedAt time.Time
}

type MessageSerializer struct{}

func (message *Message) Serialize() string {
	return fmt.Sprintf(
		"%d, %s, %d, %d, %s, %s\n",
		message.Id,
		message.Message,
		message.From,
		message.To,
		public.DatetimeToString(&message.SendAt),
		public.DatetimeToString(&message.ReceivedAt),
	)
}

func (message *Message) CompareByUser(from, to int) bool {
	if message.From == from && message.To == to {
		return true
	}
	return false
}

func (serializer *MessageSerializer) Serialize(message public.FileEntity) string {
	return message.Serialize()
}

func (serializer *MessageSerializer) Deserialize(target string) interface{} {
	data := strings.Split(target, ", ")
	if len(data) < 6 {
		public.GetWarningLogger().Printf("(%s) is invalid row.\n", target)
		return nil
	}
	return &Message{
		Id:         public.ParseInt(data[0]),
		Message:    data[1],
		From:       public.ParseInt(data[2]),
		To:         public.ParseInt(data[3]),
		SendAt:     public.ParseDatetime(data[4]),
		ReceivedAt: public.ParseDatetime(data[5]),
	}
}
