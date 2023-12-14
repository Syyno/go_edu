package auth

import (
	"context"
	"users/internal/converter"
	auth_v1 "users/pkg/auth/v1"
)

func (i *Implementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	tkn, err := i.authService.Login(ctx, converter.LoginRequestToDto(req))
	if err != nil {
		return nil, err
	}

	return &auth_v1.LoginResponse{RefreshToken: tkn}, nil
}
