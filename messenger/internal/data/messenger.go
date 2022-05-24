package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
)

type messengerRepo struct {
	data *Data
	log  *log.Helper
}

// NewMessengerRepo .
func NewMessengerRepo(data *Data, logger log.Logger) biz.MessengerRepo {
	return &messengerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *messengerRepo) FindChatByID(ctx context.Context, chatId uint64) (*biz.Chat, error) {
	return nil, nil
}

func (r *messengerRepo) FindUserByID(ctx context.Context, userId uint64) (*biz.User, error) {
	return nil, nil
}

func (r *messengerRepo) SaveMessage(ctx context.Context, msg *biz.ChatMessage) (*biz.ChatMessage, error) {
	return nil, nil
}
