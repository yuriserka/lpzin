package chats

import (
	"github.com/dustin/go-broadcast"
)

var chatChannels = make(map[uint]broadcast.Broadcaster)

func Chat(chatid uint) broadcast.Broadcaster {
	b, ok := chatChannels[chatid]
	if !ok {
		b = broadcast.NewBroadcaster(100)
		chatChannels[chatid] = b
	}
	return b
}
