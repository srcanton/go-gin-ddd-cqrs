package repository

import (
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/aggregate"
	repository "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/repository"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/persistence/model"
	"gorm.io/gorm"
)

// Repository is the repository of domain.Repository
type Repository struct {
	postgressClient *gorm.DB
}

func New(postgressClient *gorm.DB) repository.Interface {
	return &Repository{postgressClient: postgressClient}
}

func (r *Repository) Save(product *aggregate.Product) error {
	productModel := model.Product{
		ID:   product.ID.GetValue(),
		Name: product.Name.GetValue(),
		Cost: product.Cost.GetValue(),
	}
	return r.postgressClient.Create(&productModel).Error
}

func (r *Repository) GetAll() (*[]aggregate.Product, error) {
	var products []model.Product
	var aggregateProducts []aggregate.Product
	err := r.postgressClient.Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		aggregateProduct, _ := aggregate.FromRaw(product.ToArray())
		aggregateProducts = append(aggregateProducts, *aggregateProduct)
	}
	return &aggregateProducts, nil
}
