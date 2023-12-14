package sys

import (
	"errors"
	appCodes "users/internal/common/sys/code"
)

type commonError struct {
	msg  string
	code appCodes.ErrorCode
}

func NewCommonError(msg string, code appCodes.ErrorCode) *commonError {
	return &commonError{msg, code}
}

func (r *commonError) Error() string {
	return r.msg
}

func (r *commonError) Code() appCodes.ErrorCode {
	return r.code
}

func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
