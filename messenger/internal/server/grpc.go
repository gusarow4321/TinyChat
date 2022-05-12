package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/gusarow4321/TinyChat/messenger/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, messenger *service.MessengerService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterMessengerServer(srv, messenger)
	return srv
}
