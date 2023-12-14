package handlers

import (
	envconfig "users/internal/config/env"
	abstraction "users/internal/service/access"
)

type service struct {
	authConfig envconfig.AuthConfig
}

func NewAccessService(authConfig envconfig.AuthConfig) abstraction.AccessService {
	return &service{authConfig: authConfig}
}
