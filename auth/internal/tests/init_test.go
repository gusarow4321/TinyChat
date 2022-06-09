package tests

import (
	"context"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/auth/internal/conf"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"os"
)

const bufSize = 1024 * 1024

var (
	logger log.Logger
	bc     conf.Bootstrap

	configPath = "../../configs/test.yaml"
)

type mockedUserRepo struct {
	mock.Mock
}

func (m *mockedUserRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	args := m.Called(u.Name)
	return &biz.User{
		Name:     args.String(0),
		Email:    args.String(1),
		Password: args.String(2),
	}, args.Error(3)
}

func (m *mockedUserRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	args := m.Called(email)
	return &biz.User{
		Name:     args.String(0),
		Email:    args.String(1),
		Password: args.String(2),
	}, args.Error(3)
}

func init() {
	logger = log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	c.Close()
}

func newAuthClient(ctx context.Context, mocked *mockedUserRepo) (v1.AuthClient, func(), error) {
	lis := bufconn.Listen(bufSize)

	s := grpc.NewServer()
	authService, err := wireService(bc.Hasher, bc.TokenMaker, mocked, logger)
	if err != nil {
		return nil, nil, err
	}
	v1.RegisterAuthServer(s, authService)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.NewHelper(logger).Errorf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.NewHelper(logger).Errorf("Connection closing error: %v", err)
		}
		s.Stop()
	}

	return v1.NewAuthClient(conn), cleanup, nil
}
