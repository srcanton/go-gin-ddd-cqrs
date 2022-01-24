package create_user

import (
	"github.com/gin-gonic/gin"
	command "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/command/create_user"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
)

type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
}

// New create profile controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		util:       util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/register", controller.create)
}
