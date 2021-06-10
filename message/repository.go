package message

import (
	"sync"

	"simple-chat-service/public"
)

var repository *MessageRepository
var repositorySync sync.Once

type MessageRepository struct {
	public.FileRepository
}

func GetMessageRepository() *MessageRepository {
	repositorySync.Do(func() {
		repository = &MessageRepository{
			public.FileRepository{
				FileName:   "message.txt",
				Serializer: &MessageSerializer{},
			},
		}
	})
	return repository
}

func (repo *MessageRepository) GetMessages(from, to int) *[]Message {
	entities, err := repo.GetAll()

	if err != nil {
		return nil
	}

	messages := []Message{}
	for _, entity := range *entities {
		switch v := entity.(type) {
		case *Message:
			messages = append(messages, *v)
		}
	}

	targets := []Message{}
	for _, message := range messages {
		if message.CompareByUser(from, to) {
			targets = append(targets, message)
		}
	}

	return &targets
}
