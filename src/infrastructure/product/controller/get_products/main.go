package get_products

import (
	"github.com/gin-gonic/gin"
	query "github.com/srcanton/go-gin-ddd-cqrs/src/application/product/query/get_products"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
)

type Controller struct {
	api_version *gin.RouterGroup
	queryBus    *query.Bus
	util        *util.Util
}

// New create profile controller instance
func New(
	api_version *gin.RouterGroup,
	queryBus *query.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{
		api_version: api_version,
		queryBus:    queryBus,
		util:        util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.api_version.GET("/products", controller.getAll)
}
