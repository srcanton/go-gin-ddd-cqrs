package get_products

import (
	"errors"
	get_products_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/service/get_products"
)

// Bus profile query bus
type Bus struct {
	service *get_products_service.GetProducts
}

// New create bus instance
func New(service *get_products_service.GetProducts) *Bus {
	return &Bus{service: service}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*[]GetProductQueryResponse, error) {
	switch query := query.(type) {
	case *GetProductsQuery:
		return bus.handleFindProductsQuery(query)
	default:
		return nil, errors.New("invalid query")
	}
}
