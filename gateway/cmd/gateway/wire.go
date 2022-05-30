//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	conf "github.com/gusarow4321/TinyChat/gateway/internal/config"
	"github.com/gusarow4321/TinyChat/gateway/internal/interceptors"
	"github.com/gusarow4321/TinyChat/gateway/internal/registers"
	"net/http"
)

func wireApp(*conf.Rest, *conf.Auth, *conf.Messenger, *conf.Tracing, log.Logger) (*http.Server, func(), error) {
	panic(wire.Build(
		interceptors.ProviderSet,
		registers.ProviderSet,
		newTracing,
		newGatewayServer),
	)
}
