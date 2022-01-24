package login_user

import (
	email_value_object "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
	password_value_object "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/password"
)

type LoginUserQuery struct {
	Email    email_value_object.Email
	Password password_value_object.Password
}

func NewUserLoginQuery(
	Email string,
	Password string,
) (*LoginUserQuery, error) {
	userEmailValueObject, err := email_value_object.FromValue(Email)
	if err != nil {
		return nil, err
	}
	userPasswordValueObject, _ := password_value_object.FromValue(Password)
	return &LoginUserQuery{
		Email:    userEmailValueObject,
		Password: userPasswordValueObject,
	}, nil
}
