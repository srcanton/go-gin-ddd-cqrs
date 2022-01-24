package login_user

import (
	"errors"
	login_user_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/service/login_user"
)

// Bus profile query bus
type Bus struct {
	loginService *login_user_service.LoginUserService
}

// New create bus instance
func New(loginService *login_user_service.LoginUserService) *Bus {
	return &Bus{loginService: loginService}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (map[string]interface{}, error) {
	switch query := query.(type) {
	case *LoginUserQuery:
		return bus.handleLoginUserQuery(query)
	default:
		return nil, errors.New("invalid query")
	}
}
