package tests

import (
	"context"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
	"github.com/stretchr/testify/mock"
	"time"
)

type mockedMessengerRepo struct {
	mock.Mock
}

func (m *mockedMessengerRepo) FindChatByID(ctx context.Context, chatId int64) (*biz.Chat, error) {
	args := m.Called(chatId)
	return &biz.Chat{
		ID:      int64(args.Int(0)),
		OwnerID: int64(args.Int(1)),
	}, args.Error(2)
}

func (m *mockedMessengerRepo) FindUserByID(ctx context.Context, userId int64) (*biz.User, error) {
	args := m.Called(userId)
	return &biz.User{
		ID:    int64(args.Int(0)),
		Name:  args.String(1),
		Color: int32(args.Int(2)),
	}, args.Error(3)
}

func (m *mockedMessengerRepo) SaveMessage(ctx context.Context, msg *biz.ChatMessage) (*biz.ChatMessage, error) {
	args := m.Called(msg.Text)
	return &biz.ChatMessage{
		ID:        int64(args.Int(0)),
		ChatID:    int64(args.Int(1)),
		UserID:    int64(args.Int(2)),
		Text:      args.String(3),
		Timestamp: args.Get(4).(time.Time),
	}, args.Error(5)
}

type mockedProducer struct {
	mock.Mock
}

func (m *mockedProducer) Write(ctx context.Context, chatId int64, msg *v1.NewMessage) error {
	args := m.Called(chatId, msg.Text)
	return args.Error(0)
}
