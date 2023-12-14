package handlers

import (
	"context"
	"users/internal/DTO/user"
	"users/internal/converter"
)

func (s *service) Update(ctx context.Context, user user.UserUpdate) error {

	if !user.EmailProvided && !user.RoleProvided && !user.NameProvided {
		return nil
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		oldUser, err := s.userRepository.Get(ctx, user.Id)
		if err != nil {
			return err
		}

		err = s.userRepository.Update(ctx, user)
		if err != nil {
			return err
		}

		err = s.userHistoryRepository.Create(ctx, converter.ToUserHistoryModel(oldUser, &user))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
