package user_errors

import (
	"errors"
	"fmt"
)

type IncorrectUserOrPasswordError struct {
	errorMessage error
}

func (e *IncorrectUserOrPasswordError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func IncorrectUserOrPassword(errorMessage string) *IncorrectUserOrPasswordError {
	e := new(IncorrectUserOrPasswordError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
