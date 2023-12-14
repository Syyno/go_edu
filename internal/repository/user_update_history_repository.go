package repository

import (
	"context"

	"users/internal/DTO/history"
)

type UserHistoryRepository interface {
	Create(ctx context.Context, userUpdate *history.UserUpdateHistory) error
}
