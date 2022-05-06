//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"auth/internal/biz"
	"auth/internal/conf"
	"auth/internal/data"
	"auth/internal/pkg/hash"
	"auth/internal/pkg/paseto"
	"auth/internal/server"
	"auth/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Hasher, *conf.TokenMaker, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		hash.ProviderSet,
		paseto.ProviderSet,
		newApp),
	)
}
