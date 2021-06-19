package push

import (
	"time"
)

type PushMessage struct {
	Content    string
	SenderId   int64
	ReceiverId int64
	SendAt     time.Time
}
