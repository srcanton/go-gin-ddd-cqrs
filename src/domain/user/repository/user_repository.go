package user_repository

import (
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/aggregate"
	user_email "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
)

// Interface is interface of user repository
type Interface interface {
	Save(*aggregate.User) error
	GetUserByEmail(email user_email.Email) (*aggregate.User, error)
}
