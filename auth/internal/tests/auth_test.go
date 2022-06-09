package tests

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/assert/v2"
	v1 "github.com/gusarow4321/TinyChat/auth/api/auth/v1"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/hash"
	"github.com/gusarow4321/TinyChat/auth/internal/pkg/paseto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	mocked := new(mockedUserRepo)
	mocked.On("Save", "new").Return("new", "new email", "pass", nil)
	mocked.On("Save", "exists").Return("", "", "", errors.New("user already exist"))

	client, cleanup, err := newAuthClient(ctx, mocked)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	resp, err := client.SignUp(ctx, &v1.SignUpRequest{
		Name:     "new",
		Email:    "new email",
		Password: "pass",
	})
	if err != nil {
		t.Fatalf("SignUp failed: %v", err)
	}
	assert.Equal(t, "new", resp.Name)
	assert.Equal(t, "new email", resp.Email)
	assert.NotEqual(t, "", resp.Tokens.AccessToken)
	assert.NotEqual(t, "", resp.Tokens.RefreshToken)

	_, err = client.SignUp(ctx, &v1.SignUpRequest{
		Name:     "exists",
		Email:    "email",
		Password: "pass",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.AlreadyExists)
	} else {
		t.Fatal("Exists failed")
	}

	t.Cleanup(func() {
		mocked.AssertExpectations(t)
		cleanup()
	})
}

func TestSignIn(t *testing.T) {
	ctx := context.Background()

	mocked := new(mockedUserRepo)
	mocked.On("FindByEmail", "1").Return("1", "1", hash.NewPasswordHasher(bc.Hasher).Hash("pass"), nil)
	mocked.On("FindByEmail", "2").Return("2", "2", "compare err", nil)
	mocked.On("FindByEmail", "3").Return("", "", "", errors.New("not found"))

	client, cleanup, err := newAuthClient(ctx, mocked)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	resp, err := client.SignIn(ctx, &v1.SignInRequest{
		Email:    "1",
		Password: "pass",
	})
	if err != nil {
		t.Fatalf("SignIn failed: %v", err)
	}
	assert.Equal(t, "1", resp.Name)
	assert.Equal(t, "1", resp.Email)
	assert.NotEqual(t, "", resp.Tokens.AccessToken)
	assert.NotEqual(t, "", resp.Tokens.RefreshToken)

	_, err = client.SignIn(ctx, &v1.SignInRequest{
		Email:    "2",
		Password: "pass",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.PermissionDenied)
	} else {
		t.Fatal("Wrong pass failed")
	}

	_, err = client.SignIn(ctx, &v1.SignInRequest{
		Email:    "3",
		Password: "pass",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.NotFound)
	} else {
		t.Fatal("Not exists failed")
	}

	t.Cleanup(func() {
		mocked.AssertExpectations(t)
		cleanup()
	})
}

func TestRefresh(t *testing.T) {
	ctx := context.Background()

	mocked := new(mockedUserRepo)

	client, cleanup, err := newAuthClient(ctx, mocked)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	tm, err := paseto.NewPasetoMaker(bc.TokenMaker)
	if err != nil {
		t.Fatalf("New token maker err: %v", err)
	}
	token, err := tm.NewRefreshToken("1")
	if err != nil {
		t.Fatalf("New refresh token err: %v", err)
	}

	resp, err := client.RefreshToken(ctx, &v1.RefreshTokenRequest{
		RefreshToken: token,
	})
	if err != nil {
		t.Fatalf("Refresh failed: %v", err)
	}
	assert.NotEqual(t, "", resp.Tokens.AccessToken)
	assert.NotEqual(t, "", resp.Tokens.RefreshToken)

	_, err = client.RefreshToken(ctx, &v1.RefreshTokenRequest{
		RefreshToken: "wrong token",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.Unauthenticated)
	} else {
		t.Fatal("Invalid token failed")
	}

	t.Cleanup(func() {
		mocked.AssertExpectations(t)
		cleanup()
	})
}

func TestIdentity(t *testing.T) {
	ctx := context.Background()

	mocked := new(mockedUserRepo)

	client, cleanup, err := newAuthClient(ctx, mocked)
	if err != nil {
		t.Errorf("New client error: %v", err)
	}

	var userId int64 = 1

	tm, err := paseto.NewPasetoMaker(bc.TokenMaker)
	if err != nil {
		t.Fatalf("New token maker err: %v", err)
	}
	token, err := tm.NewAccessToken(fmt.Sprintf("%v", userId))
	if err != nil {
		t.Fatalf("New access token err: %v", err)
	}

	resp, err := client.Identity(ctx, &v1.IdentityRequest{
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("Identity failed: %v", err)
	}
	assert.Equal(t, userId, resp.Id)

	_, err = client.Identity(ctx, &v1.IdentityRequest{
		AccessToken: "wrong token",
	})
	if s, ok := status.FromError(err); ok {
		assert.Equal(t, s.Code(), codes.Unauthenticated)
	} else {
		t.Fatal("Invalid token failed")
	}

	t.Cleanup(func() {
		mocked.AssertExpectations(t)
		cleanup()
	})
}
