// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/gateway/internal/config"
	"github.com/gusarow4321/TinyChat/gateway/internal/interceptors"
	"github.com/gusarow4321/TinyChat/gateway/internal/registers"
	"net/http"
)

// Injectors from wire.go:

func wireApp(rest *conf.Rest, auth *conf.Auth, messenger *conf.Messenger, tracing *conf.Tracing, logger log.Logger) (*http.Server, func(), error) {
	authInterceptor, cleanup, err := interceptors.NewAuthInterceptor(auth, logger)
	if err != nil {
		return nil, nil, err
	}
	serveMux, cleanup2, err := registers.RegisterAll(auth, messenger, authInterceptor)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, err := customHandler(serveMux, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	ochttpHandler, cleanup3, err := newTracing(tracing, handler, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server := newGatewayServer(rest, ochttpHandler)
	return server, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
