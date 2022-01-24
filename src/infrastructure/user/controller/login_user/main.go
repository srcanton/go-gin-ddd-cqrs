package login_user

import (
	"github.com/gin-gonic/gin"
	query "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/query/login_user"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
)

type Controller struct {
	route    *gin.Engine
	queryBus *query.Bus
	util     *util.Util
}

// New create profile controller instance
func New(
	route *gin.Engine,
	queryBus *query.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{
		route:    route,
		queryBus: queryBus,
		util:     util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/login", controller.login)
}
