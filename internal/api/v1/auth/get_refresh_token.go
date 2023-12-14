package auth

import (
	"context"
	auth_v1 "users/pkg/auth/v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *auth_v1.GetRefreshTokenRequest) (*auth_v1.GetRefreshTokenResponse, error) {
	tkn, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &auth_v1.GetRefreshTokenResponse{RefreshToken: tkn}, nil
}
