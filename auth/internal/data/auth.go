package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/pkg/ent/user"
	"math/rand"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	// TODO: tx

	m, err := r.data.db.User.
		Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if _, err = r.data.db.UserMetadata.
		Create().
		SetName(u.Name).
		SetUser(m).
		SetColor(rand.Int31n(16777216)).
		Save(ctx); err != nil {
		return nil, err
	}

	if _, err = r.data.db.Chat.
		Create().
		SetOwner(m).
		Save(ctx); err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       m.ID,
		Name:     u.Name,
		Email:    m.Email,
		Password: m.Password,
	}, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	u, err := r.data.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	m, err := u.QueryMetadata().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       u.ID,
		Name:     m.Name,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
