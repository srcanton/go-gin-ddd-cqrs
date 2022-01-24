package auth_service

import (
	"github.com/go-redis/redis"
	auth_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/common/service/auth"
)

type RedisService struct {
	Auth   auth_service.AuthInterface
	Client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		Auth:   NewAuth(client),
		Client: client,
	}
}
