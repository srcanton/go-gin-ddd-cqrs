package get_products

import (
	"github.com/gin-gonic/gin"
	get_products_query "github.com/srcanton/go-gin-ddd-cqrs/src/application/product/query/get_products"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/errors"
	"net/http"
)

func (controller *Controller) getAll(c *gin.Context) {

	query := get_products_query.NewGetProductsQuery()
	products, err := controller.queryBus.Handle(query)
	if err != nil {
		switch err.(type) {
		case *user_errors.AlreadyExistsError:
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	}
	c.JSON(http.StatusOK, products)
}
