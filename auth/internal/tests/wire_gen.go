// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package tests

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
	"github.com/gusarow4321/TinyChat/auth/internal/conf"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"github.com/gusarow4321/TinyChat/auth/internal/service"
)

// Injectors from wire.go:

func wireService(hasher *conf.Hasher, tokenMaker *conf.TokenMaker, userRepo biz.UserRepo, logger log.Logger) (*service.AuthService, error) {
	passwordHasher := hash.NewPasswordHasher(hasher)
	pasetoTokenMaker, err := paseto.NewPasetoMaker(tokenMaker)
	if err != nil {
		return nil, err
	}
	authUsecase := biz.NewAuthUsecase(userRepo, logger, passwordHasher, pasetoTokenMaker)
	authService := service.NewAuthService(authUsecase)
	return authService, nil
}