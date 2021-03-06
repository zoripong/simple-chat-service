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

func (message *Message) CompareByUser(from, to int) bool {
	if message.From == from && message.To == to {
		return true
	}
	return false
}

type MessageSerializer struct{}

func (serializer *MessageSerializer) Serialize(entity interface{}) string {
	message, ok := entity.(Message)
	if !ok {
		public.GetErrorLogger().Printf("%s is not message.\n", entity)
		return ""
	}
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
