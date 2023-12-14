package handlers

import (
	"users/internal/repository"
	"users/internal/repository/postgres/db"
	abstraction "users/internal/service/user"
)

type service struct {
	userRepository        repository.UserRepository
	userHistoryRepository repository.UserHistoryRepository
	txManager             db.TxManager
}

func NewUserService(
	repository repository.UserRepository,
	userHistoryRepository repository.UserHistoryRepository,
	txManager db.TxManager,
) abstraction.UserService {
	return &service{
		userRepository:        repository,
		txManager:             txManager,
		userHistoryRepository: userHistoryRepository,
	}
}
