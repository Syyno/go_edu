package handlers

import (
	"context"
	"users/internal/DTO/user"
	"users/internal/converter"
	domainErrors "users/internal/domain/errors"
)

func (s *service) Create(ctx context.Context, user user.UserCreate) (int64, error) {

	if user.Password == "" || (user.Password != user.PasswordConfirm) {
		return 0, domainErrors.InequalPasswords
	}

	domainUser := converter.ToDomainFromCreate(&user)

	id, err := s.userRepository.Create(ctx, domainUser)
	if err != nil {
		return 0, err
	}
	return id, nil
}
