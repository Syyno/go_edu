package user_tests

import (
	"context"
	"testing"
	"users/internal/api/v1/user"
	serviceMocks "users/mocks/services"

	desc "users/pkg/user/v1"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	dto "users/internal/DTO/user"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	type userServiceMock func() *serviceMocks.MockUserService

	var (
		id       = gofakeit.Int64()
		password = gofakeit.BeerName()
		request  = &desc.CreateRequest{
			Name:            gofakeit.Name(),
			Email:           gofakeit.Email(),
			Password:        password,
			PasswordConfirm: password,
			Role:            0,
		}

		serviceReq = dto.UserCreate{
			Name:            request.Name,
			Email:           request.Email,
			Password:        request.Password,
			PasswordConfirm: request.PasswordConfirm,
			Role:            request.Role,
		}

		response = &desc.CreateResponse{Id: id}

		ctx = context.Background()

		methodName = "Create"

		err             = gofakeit.Error()
		failureId int64 = 0
	)

	testCases := []struct {
		name            string
		args            args
		userServiceMock userServiceMock
		wantResp        *desc.CreateResponse
		wantErr         error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: request,
			},
			userServiceMock: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, serviceReq).Return(id, nil).Once()
				return m
			},
			wantResp: response,
			wantErr:  nil,
		},
		{
			name: "failure",
			args: args{
				ctx: ctx,
				req: request,
			},
			userServiceMock: func() *serviceMocks.MockUserService {
				m := new(serviceMocks.MockUserService)
				m.On(methodName, ctx, serviceReq).Return(failureId, err).Once()
				return m
			},
			wantResp: nil,
			wantErr:  err,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			api := user.NewImplementation(tt.userServiceMock())
			gotResp, gotErr := api.Create(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.wantErr, gotErr)
			assert.Equal(t, tt.wantResp, gotResp)
		})
	}
}
