package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/pkg/ent"
	"github.com/gusarow4321/TinyChat/pkg/ent/user"
	"math/rand"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.UserRepo {
	rand.Seed(time.Now().UnixNano())

	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			r.log.WithContext(ctx).Errorf("Rolling back transaction error: %v", rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	var userModel *ent.User
	var userMetadata *ent.UserMetadata

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if userModel, err = tx.User.
			Create().
			SetEmail(u.Email).
			SetPassword(u.Password).
			Save(ctx); err != nil {
			return err
		}

		if userMetadata, err = tx.UserMetadata.
			Create().
			SetName(u.Name).
			SetUser(userModel).
			SetColor(rand.Int31n(16777216)).
			Save(ctx); err != nil {
			return err
		}

		if _, err = tx.Chat.
			Create().
			SetOwner(userModel).
			Save(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       userModel.ID,
		Name:     userMetadata.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
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
