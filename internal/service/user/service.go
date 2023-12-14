package service

import (
	"context"
	user2 "users/internal/DTO/user"
	"users/internal/domain/user"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=UserService_mock
type UserService interface {
	Get(ctx context.Context, id int64) (*user.UserDomain, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, user user2.UserCreate) (int64, error)
	Update(ctx context.Context, user user2.UserUpdate) error
}
