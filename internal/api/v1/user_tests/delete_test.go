package user_tests

import (
	"context"
	"testing"
	"users/internal/api/v1/user"
	serviceMocks "users/mocks/services"
	desc "users/pkg/user/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	type userServiceMock func() *serviceMocks.MockUserService

	var (
		methodName = "Delete"

		ctx = context.Background()

		id      = gofakeit.Int64()
		request = &desc.DeleteRequest{Id: id}

		err = gofakeit.Error()
	)

	tc := []struct {
		name                string
		wantResp            *emptypb.Empty
		wantErr             error
		args                args
		userServiceMockFunc userServiceMock
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: request,
			},
			wantResp: &emptypb.Empty{},
			wantErr:  nil,
			userServiceMockFunc: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, id).Return(nil, nil).Once()
				return m
			},
		},
		{
			name: "failure",
			args: args{
				ctx: ctx,
				req: request,
			},
			wantResp: &emptypb.Empty{},
			wantErr:  err,
			userServiceMockFunc: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, id).Return(err).Once()
				return m
			},
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			api := user.NewImplementation(tt.userServiceMockFunc())
			gotResp, gotErr := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantResp, gotResp)
			require.Equal(t, tt.wantErr, gotErr)
		})
	}
}
