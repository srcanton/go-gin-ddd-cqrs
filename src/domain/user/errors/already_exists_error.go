package user_errors

import (
	"errors"
	"fmt"
)

type AlreadyExistsError struct {
	errorMessage error
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func AlreadyExists(errorMessage string) *AlreadyExistsError {
	e := new(AlreadyExistsError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
