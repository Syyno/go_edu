package auth

import (
	"context"
	dto "users/internal/DTO/auth"
)

type AuthService interface {
	Login(ctx context.Context, creds dto.UserCredentials) (string, error)
	GetRefreshToken(ctx context.Context, refreshTkn string) (string, error)
	GetAccessToken(ctx context.Context, refreshTkn string) (string, error)
}
