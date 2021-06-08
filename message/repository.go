package message

import (
	"sync"

	"simple-chat-service/public"
)

var repository *public.FileRepository
var repositorySync sync.Once

func GetMessageRespository() *public.FileRepository {
	repositorySync.Do(func() {
		repository = &public.FileRepository{
			FileName:   "message.txt",
			Serializer: &MessageSerializer{},
		}
	})
	return repository
}
