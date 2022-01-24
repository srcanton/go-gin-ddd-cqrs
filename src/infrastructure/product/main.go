package product

import (
	"github.com/gin-gonic/gin"
	get_products_query "github.com/srcanton/go-gin-ddd-cqrs/src/application/product/query/get_products"
	get_products_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/service/get_products"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
	get_products_controller "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/product/controller/get_products"
	product_repository "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/product/repository"
	"gorm.io/gorm"
)

// Initialize initialize profile module
func Initialize(
	api_version *gin.RouterGroup, util *util.Util, postgressClient *gorm.DB,
) {
	//**********FIND PRODUCTS*************//
	productRepository := product_repository.New(postgressClient)
	getProductsService := get_products_service.New(productRepository)
	getProductsQueryBus := get_products_query.New(getProductsService)
	get_products_controller.New(api_version, getProductsQueryBus, util)

}
