package biz

import (
	"context"
	"fmt"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
)

var (
	ErrUserNotFound  = status.Errorf(codes.NotFound, "reason: %v", v1.ErrorReason_USER_NOT_FOUND.String())
	ErrWrongPassword = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ErrorReason_WRONG_PASSWORD.String())
	ErrInvalidToken  = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ErrorReason_INVALID_TOKEN.String())
)

// User is a user model.
type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByID(context.Context, uint64) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
}

// AuthUsecase is a Greeter usecase.
type AuthUsecase struct {
	repo UserRepo
	log  *log.Helper

	hasher     hash.PasswordHasher
	tokenMaker paseto.TokenMaker
}

// NewAuthUsecase new a Greeter usecase.
func NewAuthUsecase(repo UserRepo, logger log.Logger, hasher hash.PasswordHasher, tokenMaker paseto.TokenMaker) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger), hasher: hasher, tokenMaker: tokenMaker}
}

// NewTokens generate new access & refresh paseto tokens
func (uc *AuthUsecase) NewTokens(ctx context.Context, userId uint64) (*v1.Tokens, error) {
	uc.log.WithContext(ctx).Infof("NewTokens: %v", userId)

	access, err := uc.tokenMaker.NewAccessToken(fmt.Sprintf("%v", userId))
	if err != nil {
		return nil, err
	}

	refresh, err := uc.tokenMaker.NewRefreshToken(fmt.Sprintf("%v", userId))
	if err != nil {
		return nil, err
	}

	return &v1.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// CreateUser save new user to db
func (uc *AuthUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Email)
	u.Password = uc.hasher.Hash(u.Password)
	return uc.repo.Save(ctx, u)
}

// ComparePassword compare pass with the one saved in db
func (uc *AuthUsecase) ComparePassword(ctx context.Context, email, pass string) (*User, error) {
	uc.log.WithContext(ctx).Infof("ComparePassword: %v", email)

	model, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if ok := uc.hasher.Compare(model.Password, pass); !ok {
		return nil, ErrWrongPassword
	}

	return model, nil
}

// GetIdFromRefresh parse refresh token & return user id
func (uc *AuthUsecase) GetIdFromRefresh(ctx context.Context, refresh string) (uint64, error) {
	idStr, err := uc.tokenMaker.ParseRefreshToken(refresh)
	if err != nil {
		return 0, ErrInvalidToken
	}

	uc.log.WithContext(ctx).Infof("GetIdFromRefresh: " + idStr)

	return strconv.ParseUint(idStr, 10, 64)
}

// Identity identifies user
func (uc *AuthUsecase) Identity(ctx context.Context, access string) (uint64, error) {
	idStr, err := uc.tokenMaker.ParseAccessToken(access)
	if err != nil {
		return 0, ErrInvalidToken
	}

	uc.log.WithContext(ctx).Infof("Identity: " + idStr)

	return strconv.ParseUint(idStr, 10, 64)
}
