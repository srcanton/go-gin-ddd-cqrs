package product_id

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/custom-error"
)

type ProductId struct {
	value string
}

func FromValue(value string) (ProductId, error) {
	err := assertValidValue(value)
	if err != nil {
		return ProductId{}, err
	}
	return ProductId{value: value}, err
}

func assertValidValue(value string) error {
	_, err := uuid.Parse(value)
	if err != nil {
		return custom_error.InvalidArgument(fmt.Sprintf("Invalid ProductId %s, the value should a valid uuid", value))
	}
	return nil
}

func (n ProductId) GetValue() string {
	return n.value
}
