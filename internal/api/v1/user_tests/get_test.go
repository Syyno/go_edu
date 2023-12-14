package user_tests

import (
	"context"
	"users/internal/api/v1/user"

	"users/internal/domain/errors"
	domain "users/internal/domain/user"
	serviceMocks "users/mocks/services"
	desc "users/pkg/user/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	type userServiceMock func() *serviceMocks.MockUserService

	var (
		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
	)

	var (
		methodName = "Get"

		ctx = context.Background()

		req = &desc.GetRequest{Id: id}

		resp = &desc.GetResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      0,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}

		serviceRes = &domain.UserDomain{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      0,
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}

		notFoundErr = &errors.UserNotFoundError{
			Id: id,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.GetResponse
		err                 error
		userServiceMockFunc userServiceMock
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			userServiceMockFunc: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, id).Return(serviceRes, nil).Once()
				return m
			},
		},
		{
			name: "user_not_found",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, notFoundErr.Error()),
			userServiceMockFunc: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, id).Return(nil, notFoundErr).Once()
				return m
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			api := user.NewImplementation(tt.userServiceMockFunc())
			gotUser, gotErr := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, gotUser)
			require.Equal(t, tt.err, gotErr)
		})
	}
}
