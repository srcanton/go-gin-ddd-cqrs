package create_user_command

import (
	security_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/security"
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/aggregate"
	user_password "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/password"
)

func (bus *Bus) handleCreateUserCommand(
	command *CreateUserCommand,
	security security_service.Interface,
) error {
	userHashedPassword, err := security.Hash(command.Password.GetValue())
	if err != nil {
		return err
	}
	userHashedPasswordValueObject, _ := user_password.FromValue(string(userHashedPassword))

	return bus.service.Create(
		&aggregate.User{
			ID:        command.ID,
			FirstName: command.FirstName,
			LastName:  command.LastName,
			Email:     command.Email,
			Password:  userHashedPasswordValueObject,
			CreatedAt: command.CreatedAt,
			UpdatedAt: command.UpdatedAt,
		},
	)
}
