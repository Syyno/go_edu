package tests

import (
	"context"
	"testing"
	dto "users/internal/DTO/user"
	domainErrors "users/internal/domain/errors"
	domain "users/internal/domain/user"
	"users/internal/service/user/handlers"
	repoMocks "users/mocks/repositories"
	user_v1 "users/pkg/user/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	methodName := "Create"

	type args struct {
		ctx context.Context
		u   dto.UserCreate
	}

	type userRepoMock func() *repoMocks.MockUserRepository

	var (
		name            = gofakeit.Name()
		password        = gofakeit.BeerName()
		passwordConfirm = password
		email           = gofakeit.Email()
		role            = 0

		userCreateValid = dto.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            user_v1.Role(role),
		}

		ctx = context.Background()

		userDomain = domain.UserDomain{
			Name:     userCreateValid.Name,
			Email:    userCreateValid.Email,
			Password: userCreateValid.Password,
			Role:     domain.Role(userCreateValid.Role),
		}
		outputId = gofakeit.Int64()
	)

	var (
		userCreateWrongPass = dto.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: name,
			Role:            user_v1.Role(role),
		}

		errInequalPass = domainErrors.InequalPasswords

		errorId int64 = 0
	)

	tc := []struct {
		name         string
		args         args
		wantErr      error
		wantId       int64
		userRepoMock userRepoMock
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				u:   userCreateValid,
			},
			wantErr: nil,
			wantId:  outputId,
			userRepoMock: func() *repoMocks.MockUserRepository {
				m := new(repoMocks.MockUserRepository)
				m.On(methodName, ctx, userDomain).Return(outputId, nil).Once()
				return m
			},
		},
		{
			name: "error_same_password",
			args: args{
				ctx: ctx,
				u:   userCreateWrongPass,
			},
			wantErr: errInequalPass,
			wantId:  errorId,
			userRepoMock: func() *repoMocks.MockUserRepository {
				return new(repoMocks.MockUserRepository)
			},
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			service := handlers.NewUserService(tt.userRepoMock(), nil, nil)
			gotId, gotErr := service.Create(tt.args.ctx, tt.args.u)
			assert.Equal(t, tt.wantId, gotId)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
