package auth

import (
	"context"
	auth_v1 "users/pkg/auth/v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error) {
	tkn, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}
	return &auth_v1.GetAccessTokenResponse{AccessToken: tkn}, nil
}
