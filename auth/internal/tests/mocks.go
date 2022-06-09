package tests

import (
	"context"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/stretchr/testify/mock"
)

type mockedUserRepo struct {
	mock.Mock
}

func (m *mockedUserRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	args := m.Called(u.Name)
	return &biz.User{
		ID:       int64(args.Int(0)),
		Name:     args.String(1),
		Email:    args.String(2),
		Password: args.String(3),
	}, args.Error(4)
}

func (m *mockedUserRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	args := m.Called(email)
	return &biz.User{
		ID:       int64(args.Int(0)),
		Name:     args.String(1),
		Email:    args.String(2),
		Password: args.String(3),
	}, args.Error(4)
}
