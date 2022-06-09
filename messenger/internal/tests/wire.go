//go:build wireinject
// +build wireinject

package tests

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gusarow4321/TinyChat/messenger/internal/biz"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/kafka"
	"github.com/gusarow4321/TinyChat/messenger/internal/pkg/observer"
	"github.com/gusarow4321/TinyChat/messenger/internal/service"
)

func wireService(kafka.MessagesProducer, biz.MessengerRepo, observer.ChatsObserver, log.Logger) (*service.MessengerService, error) {
	panic(wire.Build(
		biz.ProviderSet,
		service.ProviderSet),
	)
}
