package login_user

import (
	"github.com/gin-gonic/gin"
	login_user_query "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/query/login_user"
	custom_error "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/custom-error"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/errors"
	"net/http"
)

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *Controller) login(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	query, err := login_user_query.NewUserLoginQuery(
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
	user, err := controller.queryBus.Handle(query)

	if err != nil {
		switch err.(type) {
		case *user_errors.IncorrectUserOrPasswordError:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}
