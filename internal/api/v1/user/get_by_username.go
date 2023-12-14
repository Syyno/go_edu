package user

import (
	"context"

	user_v1 "users/pkg/user/v1"
)

func (i *Implementation) GetByUserName(context.Context, *user_v1.GetByUserNameRequest) (*user_v1.GetResponse, error) {
	return &user_v1.GetResponse{
		Id:        0,
		Name:      "test",
		Email:     "",
		Role:      0,
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil
}
