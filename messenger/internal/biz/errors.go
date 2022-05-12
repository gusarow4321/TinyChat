package biz

import (
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_USER_NOT_FOUND.String())
	ErrChatNotFound = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_CHAT_NOT_FOUND.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
