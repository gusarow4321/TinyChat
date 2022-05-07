package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	return u, nil
}

func (r *userRepo) FindByID(context.Context, uint64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) FindByEmail(context.Context, string) (*biz.User, error) {
	return nil, nil
}
