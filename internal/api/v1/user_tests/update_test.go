package user_tests

import (
	"context"
	"testing"
	dto "users/internal/DTO/user"
	"users/internal/api/v1/user"
	serviceMocks "users/mocks/services"
	desc "users/pkg/user/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	type userServiceMock func() *serviceMocks.MockUserService

	var (
		methodName = "Update"

		ctx = context.Background()

		request = &desc.UpdateRequest{
			Id:    gofakeit.Int64(),
			Name:  wrapperspb.String(gofakeit.Name()),
			Email: wrapperspb.String(gofakeit.Email()),
			Role:  wrapperspb.Int32(0),
		}

		updateDto = dto.UserUpdate{
			Id:            request.Id,
			NameValue:     request.Name.GetValue(),
			NameProvided:  true,
			EmailValue:    request.Email.GetValue(),
			EmailProvided: true,
			RoleValue:     int(request.Role.GetValue()),
			RoleProvided:  true,
		}

		err = gofakeit.Error()
	)

	tests := []struct {
		name                string
		args                args
		wantResp            *emptypb.Empty
		wantErr             error
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
				m.On(methodName, ctx, updateDto).Return(nil).Once()
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
				m.On(methodName, ctx, updateDto).Return(err).Once()
				return m
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := user.NewImplementation(tt.userServiceMockFunc())
			gotResp, gotErr := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantResp, gotResp)
			require.Equal(t, tt.wantErr, gotErr)
		})
	}
}
