package create_user

import (
	"github.com/gin-gonic/gin"
	create_user_command "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/command/create_user"
	custom_error "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/custom-error"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/errors"
	"net/http"
)

type CreateUserInput struct {
	ID        string `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (controller *Controller) create(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	command, err := create_user_command.NewCreateUserCommand(
		input.ID,
		input.FirstName,
		input.LastName,
		input.Email,
		input.Password,
	)

	if err != nil {
		switch err.(type) {
		case *custom_error.InvalidArgumentError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	}
	err = controller.commandBus.Handle(command)
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
	c.JSON(http.StatusCreated, nil)
}
