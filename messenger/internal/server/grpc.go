package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/gusarow4321/TinyChat/messenger/internal/conf"
	"github.com/gusarow4321/TinyChat/messenger/internal/service"
	pkgmetrics "github.com/gusarow4321/TinyChat/pkg/metrics"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, messenger *service.MessengerService, vecs *pkgmetrics.Vecs, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				validate.Validator(),
				tracing.Server(),
				logging.Server(logger),
				metrics.Server(
					metrics.WithSeconds(prom.NewHistogram(vecs.Seconds)),
					metrics.WithRequests(prom.NewCounter(vecs.Requests)),
				),
			),
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
