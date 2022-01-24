package user_email

import (
	"fmt"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/custom-error"
	"net/mail"
)

type Email struct {
	value string
}

func FromValue(value string) (Email, error) {
	err := assertValidValue(value)
	if err != nil {
		return Email{}, err
	}
	return Email{value: value}, err
}

func assertValidValue(value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid Email %s, the value should a valid email", value))
	}
	return nil
}

func (n Email) GetValue() string {
	return n.value
}
