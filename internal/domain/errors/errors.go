package errors

import (
	"errors"
	"fmt"
)

type UserNotFoundError struct {
	Id int64
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("There is no user with id=%d", e.Id)
}

var InequalPasswords = errors.New("Password is empty or not match with password confirm")
