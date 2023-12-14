package repository

import (
	"context"
	"users/internal/DTO/user"
	domain "users/internal/domain/user"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*domain.UserDomain, error)
	Create(ctx context.Context, userDomain domain.UserDomain) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, userDomain user.UserUpdate) error
}
