package security_service

import (
	security_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/security"
	"golang.org/x/crypto/bcrypt"
)

type Security struct {
}

func New() security_service.Interface {
	return &Security{}
}

func (r *Security) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (r *Security) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
