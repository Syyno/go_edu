package handlers

import (
	"context"
	"users/internal/domain/user"
)

func (s *service) Get(ctx context.Context, id int64) (*user.UserDomain, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
