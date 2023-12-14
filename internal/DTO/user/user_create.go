package user

import user_v1 "users/pkg/user/v1"

type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            user_v1.Role
}
