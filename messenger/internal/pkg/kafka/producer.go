package kafka

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/segmentio/kafka-go"
	"strconv"
)

type MessagesProducer interface {
	Write(context.Context, int64, *v1.NewMessage) error
}

type Producer struct {
	w *kafka.Writer
}

func NewProducer(conf *conf.Kafka, logger log.Logger) (MessagesProducer, func()) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(conf.Addr),
		Topic:    conf.Topic,
		Balancer: kafka.CRC32Balancer{},
		Async:    true,
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the kafka producer")
		if err := w.Close(); err != nil {
			log.NewHelper(logger).Errorf("error while closing the kafka producer: %v", err)
		}
	}
	return &Producer{w: w}, cleanup
}

func (p *Producer) Write(ctx context.Context, chatId int64, msg *v1.NewMessage) error {
	v, err := kafka.Marshal(kafkaNewMsg{
		ID:        msg.Id,
		ChatID:    chatId,
		UserID:    msg.User.Id,
		Name:      msg.User.Name,
		Color:     msg.User.Color,
		Text:      msg.Text,
		Timestamp: msg.Timestamp.AsTime(),
	})
	if err != nil {
		return err
	}

	return p.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(int(chatId))),
		Value: v,
	})
}
