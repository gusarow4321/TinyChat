package tests

import (
	"context"
	"errors"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	ctx := context.Background()

	mockedRepo := new(mockedMessengerRepo)
	mockedRepo.On("FindUserByID", int64(1)).Return(0, "", 0, errors.New("user not found"))

	mockedRepo.On("FindUserByID", int64(2)).Return(2, "name", 11259375, nil)
	mockedRepo.On("FindChatByID", int64(1)).Return(0, 0, errors.New("chat not found"))

	mockedRepo.On("FindUserByID", int64(3)).Return(3, "name", 11259375, nil)
	mockedRepo.On("FindChatByID", int64(2)).Return(2, 3, nil)
	mockedRepo.On("SaveMessage", "success msg").Return(1, 2, 3, "success msg", time.Now(), nil)

	mockedProd := new(mockedProducer)
	mockedProd.On("Write", int64(2), "success msg").Return(nil)

	client, cleanup, err := newMessengerClient(ctx, mockedProd, mockedRepo)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	_, err = client.Send(ctx, &v1.SendRequest{
		ChatId: 0,
		UserId: 1,
		Text:   "",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.NotFound)
	} else {
		t.Fatal("User not found failed")
	}

	_, err = client.Send(ctx, &v1.SendRequest{
		ChatId: 1,
		UserId: 2,
		Text:   "",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.NotFound)
	} else {
		t.Fatal("Chat not found failed")
	}

	resp, err := client.Send(ctx, &v1.SendRequest{
		ChatId: 2,
		UserId: 3,
		Text:   "success msg",
	})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	assert.Equal(t, "success msg", resp.Text)
	assert.Equal(t, "#abcdef", resp.User.Color)

	t.Cleanup(func() {
		mockedRepo.AssertExpectations(t)
		mockedProd.AssertExpectations(t)
		cleanup()
	})
}

func TestChat(t *testing.T) {
	ctx := context.Background()

	mockedRepo := new(mockedMessengerRepo)
	mockedRepo.On("FindUserByID", int64(1)).Return(1, "name", 11259375, nil)
	mockedRepo.On("FindChatByID", int64(1)).Return(1, 1, nil)

	mockedProd := new(mockedProducer)

	client, cleanup, err := newMessengerClient(ctx, mockedProd, mockedRepo)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	sub, err := client.Subscribe(ctx, &v1.SubscribeRequest{
		ChatId: 1,
		UserId: 1,
	})
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	var msgId int64 = 1
	var otherId int64 = 2
	ts := timestamppb.Now()
	newMsg := &v1.NewMessage{
		Id: msgId,
		User: &v1.NewMessage_User{
			Id:    otherId,
			Name:  "other user",
			Color: "#ffffff",
		},
		Text:      "text",
		Timestamp: ts,
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				o.Publish(1, newMsg)
			}
		}
	}()

	r, err := sub.Recv()
	if err != nil {
		t.Fatalf("Receive failed: %v", err)
	}
	assert.Equal(t, msgId, r.Id)
	assert.Equal(t, otherId, r.User.Id)
	assert.Equal(t, "other user", r.User.Name)
	assert.Equal(t, "#ffffff", r.User.Color)
	assert.Equal(t, "text", r.Text)
	assert.Equal(t, ts.GetSeconds(), r.Timestamp.GetSeconds())

	done <- true

	t.Cleanup(func() {
		mockedRepo.AssertExpectations(t)
		mockedProd.AssertExpectations(t)
		cleanup()
	})
}
