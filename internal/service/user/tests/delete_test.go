package tests

import (
	"context"
	"testing"
	"users/internal/service/user/handlers"
	repoMocks "users/mocks/repositories"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	methodName := "Delete"

	type args struct {
		ctx context.Context
		id  int64
	}

	id := gofakeit.Int64()
	ctx := context.Background()

	type userRepoMock func() *repoMocks.MockUserRepository

	err := gofakeit.Error()

	tc := []struct {
		name         string
		wantErr      error
		args         args
		userRepoMock userRepoMock
	}{
		{
			name:    "success",
			wantErr: nil,
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func() *repoMocks.MockUserRepository {
				m := new(repoMocks.MockUserRepository)
				m.On(methodName, ctx, id).Return(nil).Once()
				return m
			},
		},
		{
			name:    "failure",
			wantErr: err,
			args: args{
				ctx: ctx,
				id:  id,
			},
			userRepoMock: func() *repoMocks.MockUserRepository {
				m := new(repoMocks.MockUserRepository)
				m.On(methodName, ctx, id).Return(err).Once()
				return m
			},
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := handlers.NewUserService(tt.userRepoMock(), nil, nil)
			gotErr := service.Delete(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
