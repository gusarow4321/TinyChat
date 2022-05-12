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

func (r *messengerRepo) FindChatByID(ctx context.Context, chatId uint64) (*biz.ChatInfo, error) {
	return nil, nil
}

func (r *messengerRepo) FindUserByID(ctx context.Context, userId uint64) (*biz.User, error) {
	return nil, nil
}

func (r *messengerRepo) LastMessageID(ctx context.Context, chatId uint64) (uint64, error) {
	return 0, nil
}

func (r *messengerRepo) AttachUserToChat(ctx context.Context, user *biz.User, chatId uint64) (*biz.ChatInfo, error) {
	return nil, nil
}

func (r *messengerRepo) DetachUserFromChat(ctx context.Context, userId, chatId uint64) {
	return
}

func (r *messengerRepo) SaveMessage(ctx context.Context, msg *biz.ChatMessage) (*biz.ChatMessage, error) {
	return nil, nil
}

func (r *messengerRepo) ListMessagesFrom(ctx context.Context, chatId, lastMsgId uint64, limit int32) ([]*biz.ChatMessage, error) {
	return nil, nil
}
