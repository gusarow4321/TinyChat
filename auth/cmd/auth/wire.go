//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/auth/internal/conf"
	"github.com/gusarow4321/TinyChat/auth/internal/data"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"github.com/gusarow4321/TinyChat/auth/internal/server"
	"github.com/gusarow4321/TinyChat/auth/internal/service"
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
