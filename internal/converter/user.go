package converter

import (
	"time"
	"users/internal/DTO/history"
	"users/internal/DTO/user"
	domain "users/internal/domain/user"
	api "users/pkg/user/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPresentation(user *domain.UserDomain) api.GetResponse {
	vmUser := api.GetResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      api.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	if user.UpdatedAt != nil {
		vmUser.UpdatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return vmUser
}

func ToDomain(vm *api.CreateRequest) domain.UserDomain {
	return domain.UserDomain{
		Name:      vm.GetName(),
		Email:     vm.GetEmail(),
		Role:      domain.Role(vm.Role),
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
}

func ToCreateModel(vm *api.CreateRequest) user.UserCreate {
	return user.UserCreate{
		Name:            vm.GetName(),
		Email:           vm.GetEmail(),
		Password:        vm.GetPassword(),
		PasswordConfirm: vm.GetPasswordConfirm(),
		Role:            vm.Role,
	}
}

func ToUpdateModel(vm *api.UpdateRequest) user.UserUpdate {
	updateModel := user.UserUpdate{Id: vm.GetId()}
	if vm.Name != nil {
		updateModel.NameProvided = true
		updateModel.NameValue = vm.GetName().GetValue()
	}

	if vm.Email != nil {
		updateModel.EmailProvided = true
		updateModel.EmailValue = vm.GetEmail().GetValue()
	}

	if vm.Role != nil {
		updateModel.RoleProvided = true
		updateModel.RoleValue = int(vm.GetRole().GetValue())
	}

	return updateModel
}

func ToDomainFromCreate(cm *user.UserCreate) domain.UserDomain {
	return domain.UserDomain{
		Name:      cm.Name,
		Email:     cm.Email,
		Password:  cm.Password,
		Role:      domain.Role(cm.Role),
		CreatedAt: time.Time{},
		UpdatedAt: nil,
	}
}

func ToUserHistoryModel(oldUser *domain.UserDomain, newUser *user.UserUpdate) *history.UserUpdateHistory {
	history := history.UserUpdateHistory{
		UserId:   oldUser.Id,
		EmailOld: oldUser.Email,
		NameOld:  oldUser.Name,
		RoleOld:  int(oldUser.Role),
	}

	if newUser.EmailProvided {
		history.EmailNew = newUser.EmailValue
	}
	if newUser.RoleProvided {
		history.RoleNew = newUser.RoleValue
	}
	if newUser.NameProvided {
		history.NameNew = newUser.NameValue
	}

	return &history
}
