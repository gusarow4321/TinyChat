package interceptors

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	conf "github.com/gusarow4321/TinyChat/gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

const (
	authorizationHeader = runtime.MetadataPrefix + "authorization"
)

type AuthInterceptor struct {
	client v1.AuthClient
}

func NewAuthInterceptor(conf *conf.Auth, logger log.Logger) (*AuthInterceptor, func(), error) {
	conn, err := grpc.Dial(conf.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the auth connection")
		if err := conn.Close(); err != nil {
			log.NewHelper(logger).Errorf("error while closing auth connection: %v", err)
		}
	}

	return &AuthInterceptor{v1.NewAuthClient(conn)}, cleanup, nil
}

func (i *AuthInterceptor) identity(ctx context.Context) error {
	val := metautils.ExtractOutgoing(ctx).Get(authorizationHeader)

	_, err := i.client.Identity(ctx, &v1.IdentityRequest{AccessToken: strings.TrimPrefix(val, "Bearer ")})
	if err != nil {
		return err
	}

	return nil
}

func (i *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if err := i.identity(ctx); err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (i *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		if err := i.identity(ctx); err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
