// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/gusarow4321/TinyChat/messenger/internal/data"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/kafka"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/observer"
	"github.com/gusarow4321/TinyChat/messenger/internal/server"
	"github.com/gusarow4321/TinyChat/messenger/internal/service"
	"github.com/gusarow4321/TinyChat/pkg/metrics"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confKafka *conf.Kafka, vecs *metrics.Vecs, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	messengerRepo := data.NewMessengerRepo(dataData, logger)
	chatsObserver := observer.NewObserver(logger)
	messengerUsecase := biz.NewMessengerUsecase(messengerRepo, logger, chatsObserver)
	messagesProducer, cleanup2 := kafka.NewProducer(confKafka, logger)
	messengerService := service.NewMessengerService(messengerUsecase, messagesProducer)
	grpcServer := server.NewGRPCServer(confServer, messengerService, vecs, logger)
	httpServer := server.NewHTTPServer(confServer)
	consumerServer := kafka.NewConsumerServer(confKafka, chatsObserver, logger)
	app := newApp(logger, grpcServer, httpServer, consumerServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
