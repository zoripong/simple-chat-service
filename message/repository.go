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
