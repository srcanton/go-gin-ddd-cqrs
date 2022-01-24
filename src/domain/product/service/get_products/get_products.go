package get_products_service

import (
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/aggregate"
	product_repository "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/repository"
)

// GetUserByEmail query bus
type GetProducts struct {
	repository product_repository.Interface
}

// New create bus instance
func New(repository product_repository.Interface) *GetProducts {
	return &GetProducts{repository: repository}
}

func (s *GetProducts) GetProducts() (*[]aggregate.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
