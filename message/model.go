package message

import (
	"fmt"
	"strings"
	"time"

	"go-gin-mongo/public"
)

type Message struct {
	Id int
	Message string
	From int
	To int
	SendAt time.Time
	ReceivedAt time.Time
}

type MessageSerializer struct {}

type User struct {
	Id int
	name string
}

func (message *Message) EqualsId(id int) bool {
	if message.Id == id {
		return true
	}
	return false
}

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

func (serializer *MessageSerializer) Serialize(message public.FileEntity) string {
	return message.Serialize()
}

func (serializer *MessageSerializer) Deserialize(target string) public.FileEntity {
	data := strings.Split(target, ", ")
	return &Message{
		Id: public.ParseInt(data[0]),
		Message: data[1],
		From: public.ParseInt(data[2]),
		To: public.ParseInt(data[3]),
		SendAt: public.ParseDatetime(data[4]),
		ReceivedAt: public.ParseDatetime(data[5]),
	}
}