package product_repository

import (
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/aggregate"
)

// Interface is interface of product repository
type Interface interface {
	Save(product *aggregate.Product) error
	GetAll() (*[]aggregate.Product, error)
}
