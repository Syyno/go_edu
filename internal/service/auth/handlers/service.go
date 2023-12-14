package handlers

import (
	envconfig "users/internal/config/env"
	abstraction "users/internal/service/auth"
	userService "users/internal/service/user"
)

type service struct {
	authConfig  envconfig.AuthConfig
	userService userService.UserService
}

func NewAuthService(authConfig envconfig.AuthConfig, userService userService.UserService) abstraction.AuthService {
	return &service{authConfig: authConfig, userService: userService}
}
