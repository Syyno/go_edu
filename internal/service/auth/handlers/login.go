package handlers

import (
	"context"
	"errors"
	dto "users/internal/DTO/auth"
	"users/internal/common/models/auth"
	"users/internal/common/utils"
)

func (s *service) Login(ctx context.Context, creds dto.UserCredentials) (string, error) {
	// Лезем в базу или кэш за данными пользователя
	// Сверяем хэши пароля

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: creds.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(s.authConfig.RefreshTokenSecretKey()),
		s.authConfig.RefreshTokenExpiration(),
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil

}
