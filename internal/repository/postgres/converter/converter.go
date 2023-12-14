package converter

import (
	domain "users/internal/domain/user"
	"users/internal/repository/postgres/user/model"
)

func ToDomain(repoModel *model.UserRepo) *domain.UserDomain {
	return &domain.UserDomain{
		Id:        repoModel.Id,
		Name:      repoModel.Name,
		Email:     repoModel.Email,
		Password:  repoModel.Password,
		Role:      domain.Role(repoModel.Role),
		CreatedAt: repoModel.CreatedAt,
		UpdatedAt: repoModel.UpdatedAt,
	}
}
