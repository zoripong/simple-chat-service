package message

import (
	"sync"
	"simple-chat-service/public"
)
 
var instance *public.FileRepository
var once sync.Once
 
func GetMessageRespository() *public.FileRepository {
    once.Do(func () {
        instance = &public.FileRepository{
					FileName: "message.txt",
					Serializer: &MessageSerializer{},
				}
    })
    return instance
}