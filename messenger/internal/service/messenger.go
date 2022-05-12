package service

import (
	"context"
	"fmt"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
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
	msg, err := s.uc.Send(ctx, in.ChatId, in.UserId, in.Text)
	if err != nil {
		return nil, err
	}

	return &v1.NewMessage{
		Id: msg.ID,
		User: &v1.User{
			Id:    msg.UserID,
			Name:  msg.Name,
			Color: fmt.Sprintf("%x", msg.Color),
		},
		Text: msg.Text,
	}, err
}
