package biz

import (
	"context"
	"fmt"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
)

// User is a user model.
type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
}

// AuthUsecase is an Auth usecase.
type AuthUsecase struct {
	repo UserRepo
	log  *log.Helper

	hasher     hash.PasswordHasher
	tokenMaker paseto.TokenMaker
}

// NewAuthUsecase new an Auth usecase.
func NewAuthUsecase(repo UserRepo, logger log.Logger, hasher hash.PasswordHasher, tokenMaker paseto.TokenMaker) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger), hasher: hasher, tokenMaker: tokenMaker}
}

// NewTokens generate new access & refresh paseto tokens
func (uc *AuthUsecase) NewTokens(ctx context.Context, userId int64) (*v1.Tokens, error) {
	uc.log.WithContext(ctx).Infof("NewTokens: %v", userId)

	access, err := uc.tokenMaker.NewAccessToken(fmt.Sprintf("%v", userId))
	if err != nil {
		return nil, internalErr(err)
	}

	refresh, err := uc.tokenMaker.NewRefreshToken(fmt.Sprintf("%v", userId))
	if err != nil {
		return nil, internalErr(err)
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
	saved, err := uc.repo.Save(ctx, u)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("CreateUser error: %v", err)
		return nil, ErrUserAlreadyExists
	}
	return saved, nil
}

// ComparePassword compare pass with the one saved in db
func (uc *AuthUsecase) ComparePassword(ctx context.Context, email, pass string) (*User, error) {
	uc.log.WithContext(ctx).Infof("ComparePassword: %v", email)

	model, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("ComparePassword error: %v", err)
		return nil, ErrUserNotFound
	}

	if ok := uc.hasher.Compare(model.Password, pass); !ok {
		uc.log.WithContext(ctx).Errorf("ComparePassword error: %v", err)
		return nil, ErrWrongPassword
	}

	return model, nil
}

// GetIdFromRefresh parse refresh token & return user id
func (uc *AuthUsecase) GetIdFromRefresh(ctx context.Context, refresh string) (int64, error) {
	idStr, err := uc.tokenMaker.ParseRefreshToken(refresh)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("GetIdFromRefresh error: %v", err)
		return 0, ErrInvalidToken
	}

	uc.log.WithContext(ctx).Infof("GetIdFromRefresh: " + idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, internalErr(err)
	}
	return id, nil
}

// Identity identifies user
func (uc *AuthUsecase) Identity(ctx context.Context, access string) (int64, error) {
	idStr, err := uc.tokenMaker.ParseAccessToken(access)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Identity error: %v", err)
		return 0, ErrInvalidToken
	}

	uc.log.WithContext(ctx).Infof("Identity: " + idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, internalErr(err)
	}
	return id, nil
}
