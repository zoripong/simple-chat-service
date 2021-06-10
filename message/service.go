package message

import (
	"sync"
	"time"
)

type SendMessageData struct {
	Id      int
	Message string
	From    int
	To      int
}

type MessageService struct {
	repository *MessageRepository
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

func (service *MessageService) GetMessages(from, to int) *[]Message {
	return service.repository.GetMessages(from, to)
}
