package message

import (
	"simple-chat-service/public"
	"sync"
)

var instance *public.FileRepository
var once sync.Once

func GetMessageRespository() *public.FileRepository {
	once.Do(func() {
		instance = &public.FileRepository{
			FileName:   "message.txt",
			Serializer: &MessageSerializer{},
		}
	})
	return instance
}
