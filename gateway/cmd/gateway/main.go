package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authv1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	conf "github.com/gusarow4321/TinyChat/gateway/internal/config"
	"github.com/gusarow4321/TinyChat/gateway/internal/interceptors"
	messengerv1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// create mux & register all grpc handlers
func register(ctx context.Context, authConn *grpc.ClientConn, conf *conf.Bootstrap) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	authInt := interceptors.NewAuthInterceptor(authConn)

	authOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, conf.Auth.Addr, authOpts); err != nil {
		return nil, err
	}

	messengerOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInt.Unary()),
		grpc.WithStreamInterceptor(authInt.Stream()),
	}
	if err := messengerv1.RegisterMessengerHandlerFromEndpoint(ctx, mux, conf.Messenger.Addr, messengerOpts); err != nil {
		return nil, err
	}

	return mux, nil
}

func main() {
	flag.Parse()

	// TODO: logger

	// ctx
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// config
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// auth client connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(bc.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Register gRPC server endpoints
	mux, err := register(ctx, conn, &bc)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(bc.Rest.Addr, wsproxy.WebsocketProxy(mux)); err != nil {
		panic(err)
	}
}
