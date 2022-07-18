package observer

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"sync"
)

// ProviderSet is observer providers.
var ProviderSet = wire.NewSet(NewObserver)

type ChatsObserver interface {
	Register(int64, int64, chan *v1.NewMessage)
	Deregister(int64, int64)
	Publish(int64, *v1.NewMessage)
}

type chat struct {
	channels map[int64]chan *v1.NewMessage
}

type Observer struct {
	sync.RWMutex
	chats map[int64]*chat
	log   *log.Helper
}

func NewObserver(logger log.Logger) ChatsObserver {
	return &Observer{
		chats: make(map[int64]*chat),
		log:   log.NewHelper(logger),
	}
}

func (o *Observer) Register(chatId, userId int64, channel chan *v1.NewMessage) {
	o.Lock()
	defer o.Unlock()

	c, ok := o.chats[chatId]
	if !ok {
		c = &chat{make(map[int64]chan *v1.NewMessage)}
		o.chats[chatId] = c
	}

	c.channels[userId] = channel
}

func (o *Observer) Deregister(chatId, userId int64) {
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

func (o *Observer) Publish(chatId int64, msg *v1.NewMessage) {
	o.RLock()
	defer o.RUnlock()

	c, ok := o.chats[chatId]
	if !ok {
		return
	}

	for userId, channel := range c.channels {
		if len(channel) == cap(channel) {
			o.log.Warnf("User %v didn't receive a message, the channel is full", userId)
			continue
		}

		channel <- msg
	}
}
