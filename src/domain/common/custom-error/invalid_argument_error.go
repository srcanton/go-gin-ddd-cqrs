package custom_error

import (
	"errors"
	"fmt"
)

type InvalidArgumentError struct {
	errorMessage error
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func (e *InvalidArgumentError) Make() string {
	return fmt.Sprintf("%v", e.errorMessage)
}

func InvalidArgument(errorMessage string) *InvalidArgumentError {
	e := new(InvalidArgumentError)
	e.errorMessage = errors.New(errorMessage)
	return e
}
