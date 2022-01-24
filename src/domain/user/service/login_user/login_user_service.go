package login_user_service

import (
	auth_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/auth"
	security_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/security"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/errors"
	get_user_by_email_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/service/get_user_by_email"
	user_email "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
	user_password "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/password"
)

// LoginUserService profile query bus
type LoginUserService struct {
	userFinder *get_user_by_email_service.GetUserByEmail
	security   security_service.Interface
	redis      auth_service.AuthInterface
	token      auth_service.TokenInterface
}

// New create bus instance
func New(
	userFinder *get_user_by_email_service.GetUserByEmail,
	security security_service.Interface,
	redis auth_service.AuthInterface,
	token auth_service.TokenInterface,
) *LoginUserService {
	return &LoginUserService{
		userFinder: userFinder,
		security:   security,
		redis:      redis,
		token:      token,
	}
}

func (service *LoginUserService) Login(email user_email.Email, password user_password.Password) (map[string]interface{}, error) {
	userFound, _ := service.userFinder.GetUserByEmail(email)
	if userFound != nil {
		uErr := service.security.VerifyPassword(userFound.Password.GetValue(), password.GetValue())
		if uErr != nil {
			token, tErr := service.token.CreateToken(userFound.ID.GetValue())
			if tErr != nil {
				return nil, tErr
			}
			saveErr := service.redis.CreateAuth(userFound.ID.GetValue(), token)

			if saveErr != nil {
				return nil, saveErr
			}

			userData := make(map[string]interface{})
			userData["access_token"] = token.AccessToken
			userData["refresh_token"] = token.RefreshToken
			userData["id"] = userFound.ID.GetValue()
			userData["first_name"] = userFound.FirstName.GetValue()
			userData["last_name"] = userFound.LastName.GetValue()

			return userData, nil
		}
	}
	return nil, user_errors.IncorrectUserOrPassword("Username or password does not match.")
}
