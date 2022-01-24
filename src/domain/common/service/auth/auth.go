package auth_service

import (
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/access"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/token"
)

type AuthInterface interface {
	CreateAuth(string, *token_detail.TokenDetail) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(authD *access_detail.AccessDetail) error
}
