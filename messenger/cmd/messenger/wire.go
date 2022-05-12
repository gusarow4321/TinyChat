//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/gusarow4321/TinyChat/messenger/internal/data"
	"github.com/gusarow4321/TinyChat/messenger/internal/server"
	"github.com/gusarow4321/TinyChat/messenger/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Chat, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
