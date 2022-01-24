package create_user_command

import (
	user_email "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
	user_first_name "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/first_name"
	user_id "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/id"
	user_last_name "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/last_name"
	user_password "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/password"
	"time"
)

type CreateUserCommand struct {
	ID        user_id.UserId
	FirstName user_first_name.FirstName
	LastName  user_last_name.LastName
	Email     user_email.Email
	Password  user_password.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCreateUserCommand(
	ID string,
	FirstName string,
	LastName string,
	Email string,
	Password string,
) (*CreateUserCommand, error) {
	userIdValueObject, _ := user_id.FromValue(ID)
	userFirstNameValueObject, _ := user_first_name.FromValue(FirstName)
	userLastNameValueObject, _ := user_last_name.FromValue(LastName)
	userEmailValueObject, err := user_email.FromValue(Email)
	if err != nil {
		return nil, err
	}
	userPasswordValueObject, _ := user_password.FromValue(Password)
	return &CreateUserCommand{
		ID:        userIdValueObject,
		FirstName: userFirstNameValueObject,
		LastName:  userLastNameValueObject,
		Email:     userEmailValueObject,
		Password:  userPasswordValueObject,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
