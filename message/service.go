package message

import (
	"sync"
	"time"

	"github.com/zoripong/simple-chat-service/public"
	"github.com/zoripong/simple-chat-service/push"
)

type SendMessageData struct {
	Id      int
	Message string
	From    int
	To      int
}

type MessageService struct {
	repository *MessageRepository
	pushClient *push.PushClient
}

var service *MessageService
var serviceSync sync.Once

func GetMessageService() *MessageService {
	serviceSync.Do(func() {
		service = &MessageService{
			repository: GetMessageRepository(),
		}
	})
	return service
}

func (service *MessageService) SendMessage(data *SendMessageData) {
	message := Message{
		Id:         data.Id,
		Message:    data.Message,
		From:       data.From,
		To:         data.To,
		SendAt:     time.Now(),
		ReceivedAt: time.Now(),
	}
	service.repository.Save(&message)

	pushClient := push.NewPushClient()

	// NOTE: 각각의 HTTP request는 이미 별도의 고루틴에서 돌고있기 때문에 해당 함수를 goroutine에서 처리하지 않아도 된다.
	result := pushClient.SendPushInstantly(
		&push.PushMessage{
			Content:    message.Message,
			SenderName: "보내는이",
			ReceiverId: int64(message.To),
			SendAt:     message.SendAt,
		},
	)
	public.GetInfoLogger().Printf("Push Message Reulst: %t\n", result)
}

func (service *MessageService) GetMessages(from, to int) *[]Message {
	return service.repository.GetMessages(from, to)
}
