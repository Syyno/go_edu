package user

import (
	"context"
	"errors"
	"users/internal/converter"
	appErr "users/internal/domain/errors"
	desc "users/pkg/user/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const operationNameGet = "user.get"

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req.GetId() == 0 {
		return nil, errors.New("fake error))")
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, operationNameGet)
	defer span.Finish()

	span.SetTag("id", req.GetId())

	user, err := i.userService.Get(ctx, req.GetId())

	if err != nil {
		switch err.(type) {
		case *appErr.UserNotFoundError:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, err
		}
	}

	a := converter.ToPresentation(user)
	return &a, nil
}
