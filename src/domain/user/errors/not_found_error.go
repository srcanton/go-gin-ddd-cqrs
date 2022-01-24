package user_errors

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	errorMessage error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func NotFound(errorMessage string) *NotFoundError {
	e := new(NotFoundError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
