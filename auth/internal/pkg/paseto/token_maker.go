package paseto

import (
	"encoding/hex"
	"github.com/gusarow4321/TinyChat/auth/internal/conf"
	"github.com/vk-rv/pvx"
	"time"
)

type TokenMaker interface {
	NewAccessToken(userId string) (string, error)
	NewRefreshToken(userId string) (string, error)
	ParseAccessToken(token string) (string, error)
	ParseRefreshToken(token string) (string, error)
}

type Maker struct {
	paseto        *pvx.ProtoV4Local
	accessSymKey  *pvx.SymKey
	accessTTL     time.Duration
	refreshSymKey *pvx.SymKey
	refreshTTL    time.Duration
	assertion     string
}

func NewPasetoMaker(c *conf.TokenMaker) (TokenMaker, error) {
	ak, err := hex.DecodeString(c.AccessKey)
	if err != nil {
		return nil, err
	}

	rk, err := hex.DecodeString(c.RefreshKey)
	if err != nil {
		return nil, err
	}

	return &Maker{
		paseto:        pvx.NewPV4Local(),
		accessSymKey:  pvx.NewSymmetricKey(ak, pvx.Version4),
		accessTTL:     c.AccessTtl.AsDuration(),
		refreshSymKey: pvx.NewSymmetricKey(rk, pvx.Version4),
		refreshTTL:    c.RefreshTtl.AsDuration(),
		assertion:     c.Assert,
	}, nil
}

func (m *Maker) NewAccessToken(userId string) (string, error) {
	exp := time.Now().Add(m.accessTTL)

	return m.paseto.Encrypt(
		m.accessSymKey,
		&pvx.RegisteredClaims{Subject: userId, Expiration: &exp},
		pvx.WithAssert([]byte(m.assertion)),
	)
}

func (m *Maker) NewRefreshToken(userId string) (string, error) {
	exp := time.Now().Add(m.refreshTTL)

	return m.paseto.Encrypt(
		m.refreshSymKey,
		&pvx.RegisteredClaims{Subject: userId, Expiration: &exp},
		pvx.WithAssert([]byte(m.assertion)),
	)
}

func (m *Maker) ParseAccessToken(token string) (string, error) {
	var c pvx.RegisteredClaims

	err := m.paseto.Decrypt(token, m.accessSymKey, pvx.WithAssert([]byte(m.assertion))).ScanClaims(&c)
	if err != nil {
		return "", err
	}

	return c.Subject, nil
}

func (m *Maker) ParseRefreshToken(token string) (string, error) {
	var c pvx.RegisteredClaims

	err := m.paseto.Decrypt(token, m.refreshSymKey, pvx.WithAssert([]byte(m.assertion))).ScanClaims(&c)
	if err != nil {
		return "", err
	}

	return c.Subject, nil
}
