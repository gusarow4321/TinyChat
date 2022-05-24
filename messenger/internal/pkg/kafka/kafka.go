package kafka

import (
	"github.com/google/wire"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// ProviderSet is kafka providers.
var ProviderSet = wire.NewSet(NewProducer, NewConsumerServer)

type kafkaNewMsg struct {
	ID        uint64
	ChatID    uint64
	UserID    uint64
	Name      string
	Color     string
	Text      string
	Timestamp time.Time
}

func (m kafkaNewMsg) toApiMsg() *v1.NewMessage {
	return &v1.NewMessage{
		Id: m.ID,
		User: &v1.NewMessage_User{
			Id:    m.UserID,
			Name:  m.Name,
			Color: m.Color,
		},
		Text:      m.Text,
		Timestamp: timestamppb.New(m.Timestamp),
	}
}
