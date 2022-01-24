package security_service

// Interface is interface of user repository
type Interface interface {
	Hash(password string) ([]byte, error)
	VerifyPassword(hashedPassword, password string) error
}
