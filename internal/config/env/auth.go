package envconfig

import (
	"errors"
	"time"
)

const (
	accessTokenSecretKeyEnvName   = "ACCESS_TOKEN_SECRET_KEY"
	refreshTokenSecretKeyEnvName  = "REFRESH_TOKEN_SECRET_KEY"
	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpirationEnvName  = "ACCESS_TOKEN_EXPIRATION"
)

type AuthConfig interface {
	AccessTokenSecretKey() string
	RefreshTokenSecretKey() string
	AccessTokenExpiration() time.Duration
	RefreshTokenExpiration() time.Duration
}

type authConfig struct {
	accessTokenSecretKey   string
	refreshTokenSecretKey  string
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewAuthConfig(conf *Configuration) (AuthConfig, error) {
	accessTokenSecretKey, err := conf.Get(accessTokenSecretKeyEnvName)
	if err != nil {
		return nil, errors.New("access token secret key not found")
	}

	refreshTokenSecretKey, err := conf.Get(refreshTokenSecretKeyEnvName)
	if err != nil {
		return nil, errors.New("refresh token secret key not found")
	}

	accessTokenExpiration, err := conf.Get(accessTokenExpirationEnvName)
	if err != nil {
		return nil, errors.New("access token expiration not found")
	}

	accessTokenExpirationDuration, err := time.ParseDuration(accessTokenExpiration)
	if err != nil {
		return nil, errors.New("access token expiration duration is not valid")
	}

	refreshTokenExpiration, err := conf.Get(refreshTokenExpirationEnvName)
	if err != nil {
		return nil, errors.New("refresh token expiration not found")
	}

	refreshTokenExpirationDuration, err := time.ParseDuration(refreshTokenExpiration)
	if err != nil {
		return nil, errors.New("refresh token expiration duration is not valid")
	}

	return &authConfig{
		accessTokenSecretKey:   accessTokenSecretKey,
		refreshTokenSecretKey:  refreshTokenSecretKey,
		refreshTokenExpiration: refreshTokenExpirationDuration,
		accessTokenExpiration:  accessTokenExpirationDuration,
	}, nil
}

func (a authConfig) AccessTokenSecretKey() string {
	return a.accessTokenSecretKey
}

func (a authConfig) RefreshTokenSecretKey() string {
	return a.refreshTokenSecretKey
}

func (a authConfig) AccessTokenExpiration() time.Duration {
	return a.accessTokenExpiration
}

func (a authConfig) RefreshTokenExpiration() time.Duration {
	return a.refreshTokenExpiration
}
