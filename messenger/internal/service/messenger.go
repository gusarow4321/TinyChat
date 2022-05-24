package service

import (
	"context"
	"fmt"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MessengerService is a messenger service.
type MessengerService struct {
	v1.UnimplementedMessengerServer

	uc *biz.MessengerUsecase
}

// NewMessengerService new a messenger service.
func NewMessengerService(uc *biz.MessengerUsecase) *MessengerService {
	return &MessengerService{uc: uc}
}

// Subscribe implements messenger.Subscribe.
func (s *MessengerService) Subscribe(req *v1.SubscribeRequest, conn v1.Messenger_SubscribeServer) error {
	return s.uc.Chat(req, conn)
}

// Send implements messenger.Send.
func (s *MessengerService) Send(ctx context.Context, in *v1.SendRequest) (*v1.NewMessage, error) {
	ts := timestamppb.Now()

	msg, user, err := s.uc.Send(ctx, in.ChatId, in.UserId, in.Text, ts.AsTime())
	if err != nil {
		return nil, err
	}

	// TODO: kafka

	return &v1.NewMessage{
		Id: msg.ID,
		User: &v1.NewMessage_User{
			Id:    user.ID,
			Name:  user.Name,
			Color: fmt.Sprintf("%x", user.Color),
		},
		Text:      msg.Text,
		Timestamp: ts,
	}, err
}
