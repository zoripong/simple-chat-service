package message

import (
	"sync"
	"time"

	"simple-chat-service/public"
)

type SendMessageData struct {
	Id	int
	Message	string
	From int
	To int
}

type MessageService struct {
	repository *public.FileRepository
}

var service *MessageService
var serviceSync sync.Once

func GetMessageService() *MessageService {
	serviceSync.Do(func() {
		service = &MessageService{
			repository: GetMessageRespository(),
		}
	})
	return service
}

func (service *MessageService) SendMessage(data *SendMessageData) {
	service.repository.Save(
		&Message{
			Id:         data.Id,
			Message:    data.Message,
			From:       data.From,
			To:         data.To,
			SendAt:     time.Now(),
			ReceivedAt: time.Now(),
		},
	)
}
