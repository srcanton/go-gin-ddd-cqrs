package auth_service

import (
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/access"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/valueobject/token"
	"net/http"
)

type TokenInterface interface {
	CreateToken(userid string) (*token_detail.TokenDetail, error)
	ExtractTokenMetadata(*http.Request) (*access_detail.AccessDetail, error)
}
