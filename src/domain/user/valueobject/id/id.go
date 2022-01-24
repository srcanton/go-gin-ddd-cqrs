package user_id

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/custom-error"
)

type UserId struct {
	value string
}

func FromValue(value string) (UserId, error) {
	err := assertValidValue(value)
	if err != nil {
		return UserId{}, err
	}
	return UserId{value: value}, err
}

func assertValidValue(value string) error {
	_, err := uuid.Parse(value)
	if err != nil {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid UserId %s, the value should a valid uuid", value))
	}
	return nil
}

func (n UserId) GetValue() string {
	return n.value
}
