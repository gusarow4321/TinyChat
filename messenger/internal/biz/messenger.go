package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/observer"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"
)

// User is a user model.
type User struct {
	ID    int64
	Name  string
	Color int32 // hex color, max: 16777215
}

// Chat model
type Chat struct {
	ID      int64
	OwnerID int64
}

// ChatMessage is a chat message model
type ChatMessage struct {
	ID        int64
	ChatID    int64
	UserID    int64
	Text      string
	Timestamp time.Time
}

// MessengerRepo is a Messenger repo.
type MessengerRepo interface {
	FindChatByID(context.Context, int64) (*Chat, error)
	FindUserByID(context.Context, int64) (*User, error)
	SaveMessage(context.Context, *ChatMessage) (*ChatMessage, error)
}

// MessengerUsecase is a Messenger usecase.
type MessengerUsecase struct {
	repo     MessengerRepo
	log      *log.Helper
	observer observer.ChatsObserver
}

// NewMessengerUsecase new a Messenger usecase.
func NewMessengerUsecase(repo MessengerRepo, logger log.Logger, observer observer.ChatsObserver) *MessengerUsecase {
	return &MessengerUsecase{
		repo:     repo,
		log:      log.NewHelper(logger),
		observer: observer,
	}
}

func (uc *MessengerUsecase) extractUserID(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, internalErr(errors.New("metadata not found"))
	}
	mdValue := md.Get("user-id")
	if len(mdValue) == 0 {
		return 0, internalErr(errors.New("empty metadata"))
	}
	userId, err := strconv.ParseInt(mdValue[0], 10, 64)
	if err != nil {
		return 0, internalErr(errors.New("invalid metadata"))
	}
	return userId, nil
}

func (uc *MessengerUsecase) Chat(subReq *v1.SubscribeRequest, conn v1.Messenger_SubscribeServer) error {
	userId, err := uc.extractUserID(conn.Context())
	if err != nil {
		return err
	}

	_, err = uc.repo.FindUserByID(conn.Context(), userId)
	if err != nil {
		return ErrUserNotFound
	}

	_, err = uc.repo.FindChatByID(conn.Context(), subReq.ChatId)
	if err != nil {
		return ErrChatNotFound
	}

	channel := make(chan *v1.NewMessage, 5)
	uc.observer.Register(subReq.ChatId, userId, channel)
	defer uc.observer.Deregister(subReq.ChatId, userId)

	uc.log.WithContext(conn.Context()).Infof("Connected to chat. UserID: %v, ChatID: %v", userId, subReq.ChatId)

	for {
		select {
		case <-conn.Context().Done():
			return internalErr(conn.Context().Err())
		case msg := <-channel:
			err = conn.Send(msg)
			if err != nil {
				return internalErr(err)
			}
		case <-time.After(6 * time.Hour):
			return nil
		}
	}
}

func (uc *MessengerUsecase) Send(ctx context.Context, chatId int64, text string, ts time.Time) (*ChatMessage, *User, error) {
	userId, err := uc.extractUserID(ctx)
	if err != nil {
		return nil, nil, err
	}

	user, err := uc.repo.FindUserByID(ctx, userId)
	if err != nil {
		return nil, nil, ErrUserNotFound
	}

	_, err = uc.repo.FindChatByID(ctx, chatId) // TODO: chat info
	if err != nil {
		return nil, nil, ErrChatNotFound
	}

	uc.log.WithContext(ctx).Infof("Message send. UserID: %v, ChatID: %v", userId, chatId)

	msg, err := uc.repo.SaveMessage(ctx, &ChatMessage{
		ChatID:    chatId,
		UserID:    userId,
		Text:      text,
		Timestamp: ts,
	})
	if err != nil {
		return nil, nil, internalErr(err)
	}

	return msg, user, nil
}
