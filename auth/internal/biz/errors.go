package biz

import (
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound  = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_USER_NOT_FOUND.String())
	ErrWrongPassword = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ErrorReason_WRONG_PASSWORD.String())
	ErrInvalidToken  = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ErrorReason_INVALID_TOKEN.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
