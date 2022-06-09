package service

import (
	"context"

	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	"github.com/gusarow4321/TinyChat/auth/internal/biz"
)

// AuthService is an auth service.
type AuthService struct {
	v1.UnimplementedAuthServer

	uc *biz.AuthUsecase
}

// NewAuthService new an auth service.
func NewAuthService(uc *biz.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

// SignUp implements auth.SignUp.
func (s *AuthService) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpReply, error) {
	model, err := s.uc.CreateUser(ctx, &biz.User{Name: in.Name, Email: in.Email, Password: in.Password})
	if err != nil {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return &v1.SignUpReply{
		Id:     model.ID,
		Name:   model.Name,
		Email:  model.Email,
		Tokens: tokens,
	}, nil
}

// SignIn implements auth.SignIn.
func (s *AuthService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	model, err := s.uc.ComparePassword(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return &v1.SignInReply{
		Id:     model.ID,
		Name:   model.Name,
		Email:  model.Email,
		Tokens: tokens,
	}, nil
}

// RefreshToken implements auth.RefreshToken.
func (s *AuthService) RefreshToken(ctx context.Context, in *v1.RefreshTokenRequest) (*v1.RefreshTokenReply, error) {
	userId, err := s.uc.GetIdFromRefresh(ctx, in.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokens, err := s.uc.NewTokens(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.RefreshTokenReply{
		Tokens: tokens,
	}, nil
}

// Identity implements auth.Identity.
func (s *AuthService) Identity(ctx context.Context, in *v1.IdentityRequest) (*v1.IdentityReply, error) {
	userId, err := s.uc.Identity(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	return &v1.IdentityReply{
		Id: userId,
	}, nil
}
