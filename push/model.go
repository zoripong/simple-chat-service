package push

import (
	"time"
)

type PushMessage struct {
	Content    string
	SenderName string
	ReceiverId int64
	SendAt     time.Time
}
