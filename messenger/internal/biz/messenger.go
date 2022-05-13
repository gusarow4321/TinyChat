package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"time"
)

// User is a user model.
type User struct {
	ID    uint64
	Name  string
	Color uint32 // hex color, max: 16777215
}

// ChatInfo result struct
type ChatInfo struct {
	Count   uint64
	Members []*User
}

// ChatMessage is a chat message model
type ChatMessage struct {
	ID     uint64
	ChatID uint64
	UserID uint64
	Name   string
	Color  uint32
	Text   string
}

// MessengerRepo is a Messenger repo.
type MessengerRepo interface {
	FindChatByID(context.Context, uint64) (*ChatInfo, error) // TODO: chat params
	FindUserByID(context.Context, uint64) (*User, error)
	LastMessageID(context.Context, uint64) (uint64, error)
	AttachUserToChat(context.Context, *User, uint64) (*ChatInfo, error)
	DetachUserFromChat(context.Context, uint64, uint64)
	SaveMessage(context.Context, *ChatMessage) (*ChatMessage, error)
	ListMessagesFrom(context.Context, uint64, uint64, int32) ([]*ChatMessage, error)
}

// MessengerUsecase is a Messenger usecase.
type MessengerUsecase struct {
	repo MessengerRepo
	log  *log.Helper

	limit       int32
	updDuration time.Duration // 100 * time.Millisecond
}

// NewMessengerUsecase new a Messenger usecase.
func NewMessengerUsecase(repo MessengerRepo, logger log.Logger, chat *conf.Chat) *MessengerUsecase {
	return &MessengerUsecase{
		repo:        repo,
		log:         log.NewHelper(logger),
		limit:       chat.Limit,
		updDuration: chat.Duration.AsDuration(),
	}
}

func (uc *MessengerUsecase) Chat(subReq *v1.SubscribeRequest, conn v1.Messenger_SubscribeServer) error {
	user, err := uc.repo.FindUserByID(conn.Context(), subReq.UserId)
	if err != nil {
		return ErrUserNotFound
	}

	chatInfo, err := uc.repo.AttachUserToChat(conn.Context(), user, subReq.ChatId)
	if err != nil {
		return ErrChatNotFound
	}
	defer uc.repo.DetachUserFromChat(conn.Context(), subReq.UserId, subReq.ChatId)

	lastMsg, err := uc.repo.LastMessageID(conn.Context(), subReq.ChatId)
	if err != nil {
		return internalErr(err)
	}

	t := time.NewTicker(uc.updDuration)

	for {
		msgs, err := uc.repo.ListMessagesFrom(conn.Context(), subReq.ChatId, lastMsg, uc.limit)
		if err != nil {
			return internalErr(err)
		}

		if len(msgs) == 0 {
			<-t.C
			continue
		}
		lastMsg = msgs[len(msgs)-1].ID

		chatInfo, err = uc.repo.FindChatByID(conn.Context(), subReq.ChatId)
		if err != nil {
			return internalErr(err)
		}

		err = conn.Send(&v1.SubscribeReply{
			Messages: msgModelsToReply(msgs),
			Info: &v1.SubscribeReply_ChatInfo{
				Count:   chatInfo.Count,
				Members: userModelsToReply(chatInfo.Members),
			},
		})
		if err != nil {
			return internalErr(err)
		}

		<-t.C
	}
}

func (uc *MessengerUsecase) Send(ctx context.Context, chatId, userId uint64, text string) (*ChatMessage, error) {
	user, err := uc.repo.FindUserByID(ctx, userId)
	if err != nil {
		return nil, ErrUserNotFound
	}

	_, err = uc.repo.AttachUserToChat(ctx, user, chatId)
	if err != nil {
		return nil, ErrChatNotFound
	}

	msg, err := uc.repo.SaveMessage(ctx, &ChatMessage{
		ChatID: chatId,
		UserID: userId,
		Text:   text,
	})
	if err != nil {
		return nil, internalErr(err)
	}

	return msg, nil
}
