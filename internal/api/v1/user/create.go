package user

import (
	"context"
	"users/internal/converter"
	vm "users/pkg/user/v1"
)

func (i *Implementation) Create(ctx context.Context, request *vm.CreateRequest) (*vm.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToCreateModel(request))
	if err != nil {
		return nil, err
	}

	return &vm.CreateResponse{Id: id}, err
}
