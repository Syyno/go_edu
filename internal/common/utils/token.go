package utils

import (
	"time"
	model2 "users/internal/common/models/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func GenerateToken(info model2.UserInfo, secretKey []byte, duration time.Duration) (string, error) {
	claims := model2.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: info.Username,
		Role:     info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*model2.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model2.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model2.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
