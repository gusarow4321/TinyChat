package observer

import (
	"github.com/google/wire"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"sync"
)

// ProviderSet is observer providers.
var ProviderSet = wire.NewSet(NewObserver)

type ChatsObserver interface {
	Register(uint64, uint64, chan *v1.NewMessage)
	Deregister(uint64, uint64)
	Publish(uint64, *v1.NewMessage)
}

type chat struct {
	channels map[uint64]chan *v1.NewMessage
}

type Observer struct {
	sync.Mutex
	chats map[uint64]*chat
}

func NewObserver() ChatsObserver {
	return &Observer{
		chats: make(map[uint64]*chat),
	}
}

func (o *Observer) Register(chatId, userId uint64, channel chan *v1.NewMessage) {
	o.Lock()
	defer o.Unlock()

	c, ok := o.chats[chatId]
	if !ok {
		c = &chat{make(map[uint64]chan *v1.NewMessage)}
		o.chats[chatId] = c
	}

	c.channels[userId] = channel
}

func (o *Observer) Deregister(chatId, userId uint64) {
	o.Lock()
	defer o.Unlock()

	c, ok := o.chats[chatId]
	if !ok {
		return
	}

	delete(c.channels, userId)

	if len(c.channels) == 0 {
		delete(o.chats, chatId)
	}
}

func (o *Observer) Publish(chatId uint64, msg *v1.NewMessage) {
	o.Lock()
	defer o.Unlock()

	c, ok := o.chats[chatId]
	if !ok {
		return
	}

	for _, channel := range c.channels {
		channel <- msg
	}
}
