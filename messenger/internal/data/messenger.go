package data

import (
	"context"
	"github.com/gusarow4321/TinyChat/pkg/ent/usermetadata"

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

func (r *messengerRepo) FindChatByID(ctx context.Context, chatId int64) (*biz.Chat, error) {
	c, err := r.data.db.Chat.Get(ctx, chatId)
	return &biz.Chat{
		ID:      c.ID,
		OwnerID: c.OwnerID,
	}, err
}

func (r *messengerRepo) FindUserByID(ctx context.Context, userId int64) (*biz.User, error) {
	m, err := r.data.db.UserMetadata.Query().Where(usermetadata.UserID(userId)).Only(ctx)
	return &biz.User{
		ID:    m.ID,
		Name:  m.Name,
		Color: m.Color,
	}, err
}

func (r *messengerRepo) SaveMessage(ctx context.Context, msg *biz.ChatMessage) (*biz.ChatMessage, error) {
	m, err := r.data.db.Message.
		Create().
		SetChatID(msg.ChatID).
		SetUserID(msg.UserID).
		SetText(msg.Text).
		SetTimestamp(msg.Timestamp).
		Save(ctx)
	return &biz.ChatMessage{
		ID:        m.ID,
		ChatID:    m.ChatID,
		UserID:    m.UserID,
		Text:      m.Text,
		Timestamp: m.Timestamp,
	}, err
}
