package kafka

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/observer"
	"github.com/segmentio/kafka-go"
)

type ConsumerServer struct {
	r        *kafka.Reader
	observer observer.ChatsObserver
	log      *log.Helper
}

func NewConsumerServer(conf *conf.Kafka, obs observer.ChatsObserver, logger log.Logger) *ConsumerServer {
	return &ConsumerServer{
		r: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{conf.Addr},
			GroupID: conf.GroupId,
			Topic:   conf.Topic,
		}),
		observer: obs,
		log:      log.NewHelper(logger),
	}
}

func (s *ConsumerServer) Start(ctx context.Context) error {
	for {
		kafkaMsg, err := s.r.ReadMessage(ctx)
		if err != nil {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		var msg kafkaNewMsg
		if err := kafka.Unmarshal(kafkaMsg.Value, &msg); err != nil {
			return err
		}

		s.log.WithContext(ctx).Infof("Read kafka msg. ChatId: %v MsgText: %v", msg.ChatID, msg.Text)

		s.observer.Publish(msg.ChatID, msg.toApiMsg())
	}
}

func (s *ConsumerServer) Stop(ctx context.Context) error {
	if err := s.r.Close(); err != nil {
		return err
	}
	s.log.WithContext(ctx).Info("KafkaConsumer server stopping")
	return nil
}
