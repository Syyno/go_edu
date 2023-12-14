package handlers

import (
	"context"
	"users/internal/common/models/auth"
	"users/internal/common/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRefreshToken(ctx context.Context, refreshTkn string) (string, error) {
	claims, err := utils.VerifyToken(refreshTkn, []byte(s.authConfig.RefreshTokenSecretKey()))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(s.authConfig.RefreshTokenSecretKey()),
		s.authConfig.RefreshTokenExpiration(),
	)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
