package tests

import (
	"context"
	"testing"
	domain "users/internal/domain/user"
	"users/internal/service/user/handlers"
	repoMocks "users/mocks/repositories"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()
	methodName := "Get"

	type args struct {
		ctx context.Context
		id  int64
	}

	id := gofakeit.Int64()
	ctx := context.Background()

	type userRepoMock func() *repoMocks.MockUserRepository

	updatedAt := gofakeit.Date()
	user := &domain.UserDomain{
		Id:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Name(),
		Password:  gofakeit.BeerStyle(),
		Role:      0,
		CreatedAt: gofakeit.Date(),
		UpdatedAt: &updatedAt,
	}

	err := gofakeit.Error()

	tc := []struct {
		name         string
		wantUser     *domain.UserDomain
		wantErr      error
		args         args
		userRepoMock userRepoMock
	}{
		{
			name:     "success",
			wantUser: user,
			wantErr:  nil,
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func() *repoMocks.MockUserRepository {
				m := new(repoMocks.MockUserRepository)
				m.On(methodName, ctx, id).Return(user, nil).Once()
				return m
			},
		},
		{
			name:     "error",
			wantUser: nil,
			wantErr:  err,
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func() *repoMocks.MockUserRepository {
				m := new(repoMocks.MockUserRepository)
				m.On(methodName, ctx, id).Return(nil, err).Once()
				return m
			},
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := handlers.NewUserService(tt.userRepoMock(), nil, nil)
			gotUser, gotErr := service.Get(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantUser, gotUser)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
