package message

import (
	"sync"

	"github.com/zoripong/simple-chat-service/public"
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
	errorLogger := public.GetErrorLogger()
	entities, err := repo.GetAll()
	if err != nil {
		errorLogger.Printf("Cannot read message files: %s\n", err)
		return nil
	}
	targets := []Message{}
	for _, entity := range *entities {
		switch v := entity.(type) {
		case *Message:
			if v.CompareByUser(from, to) {
				targets = append(targets, *v)
			}
		default:
			errorLogger.Printf("%s is not message", v)
		}
	}
	return &targets
}
