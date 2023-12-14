package handlers

import (
	"context"
	"errors"
	"users/internal/common/models/auth"
	"users/internal/common/utils"
)

func (s *service) GetAccessToken(ctx context.Context, refreshTkn string) (string, error) {
	claims, err := utils.VerifyToken(refreshTkn, []byte(s.authConfig.RefreshTokenSecretKey()))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	//user, err := s.userService.Get(ctx, claims.Username)

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(s.authConfig.AccessTokenSecretKey()),
		s.authConfig.AccessTokenExpiration(),
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
