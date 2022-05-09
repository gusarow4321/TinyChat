package interceptors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	"google.golang.org/grpc"
)

const (
	authorizationHeader = runtime.MetadataPrefix + "authorization"
)

type AuthInterceptor struct {
	client v1.AuthClient
}

func NewAuthInterceptor(conn *grpc.ClientConn) *AuthInterceptor {
	return &AuthInterceptor{v1.NewAuthClient(conn)}
}

func (i *AuthInterceptor) identity(ctx context.Context) error {
	val := metautils.ExtractOutgoing(ctx).Get(authorizationHeader)

	_, err := i.client.Identity(ctx, &v1.IdentityRequest{AccessToken: val})
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