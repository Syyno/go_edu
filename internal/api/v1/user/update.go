package user

import (
	"context"
	"users/internal/converter"
	vm "users/pkg/user/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Update(ctx context.Context, request *vm.UpdateRequest) (*emptypb.Empty, error) {
	err := s.userService.Update(ctx, converter.ToUpdateModel(request))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
