package registers

import (
	"context"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authv1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	conf "github.com/gusarow4321/TinyChat/gateway/internal/config"
	"github.com/gusarow4321/TinyChat/gateway/internal/interceptors"
	messengerv1 "github.com/gusarow4321/TinyChat/messenger/api/messenger/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProviderSet = wire.NewSet(RegisterAll)

func RegisterAll(auth *conf.Auth, messenger *conf.Messenger, authInt *interceptors.AuthInterceptor) (*runtime.ServeMux, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	cleanup := func() {
		cancel()
	}

	mux := runtime.NewServeMux()

	if err := registerAuth(ctx, mux, auth); err != nil {
		return nil, cleanup, err
	}
	if err := registerMessenger(ctx, mux, messenger, authInt); err != nil {
		return nil, cleanup, err
	}

	return mux, cleanup, nil
}

func registerAuth(ctx context.Context, mux *runtime.ServeMux, conf *conf.Auth) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return authv1.RegisterAuthHandlerFromEndpoint(ctx, mux, conf.Addr, opts)
}

func registerMessenger(ctx context.Context, mux *runtime.ServeMux, conf *conf.Messenger, authInt *interceptors.AuthInterceptor) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInt.Unary()),
		grpc.WithStreamInterceptor(authInt.Stream()),
	}
	return messengerv1.RegisterMessengerHandlerFromEndpoint(ctx, mux, conf.Addr, opts)
}
