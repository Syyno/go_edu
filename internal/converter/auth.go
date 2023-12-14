package converter

import (
	dto "users/internal/DTO/auth"
	api "users/pkg/auth/v1"
)

func LoginRequestToDto(req *api.LoginRequest) dto.UserCredentials {
	return dto.UserCredentials{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
}
