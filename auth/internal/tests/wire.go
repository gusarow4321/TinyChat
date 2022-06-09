//go:build wireinject
// +build wireinject

package tests

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/auth/internal/conf"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"github.com/gusarow4321/TinyChat/auth/internal/service"
)

func wireService(*conf.Hasher, *conf.TokenMaker, biz.UserRepo, log.Logger) (*service.AuthService, error) {
	panic(wire.Build(
		biz.ProviderSet,
		service.ProviderSet,
		hash.ProviderSet,
		paseto.ProviderSet),
	)
}
