package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	create_user_command "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/command/create_user"
	login_user_query "github.com/srcanton/go-gin-ddd-cqrs/src/application/user/query/login_user"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/service/create_user"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/service/get_user_by_email"
	login_user_service "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/service/login_user"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/service/auth"
	security_service "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/service/security"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
	create_user_controller "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/user/controller/create_user"
	login_user_controller "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/user/controller/login_user"

	user_repository "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/user/repository"
	"gorm.io/gorm"
)

// Initialize initialize profile module
func Initialize(
	engine *gin.Engine, util *util.Util, postgressClient *gorm.DB, redisClient *redis.Client,
) {
	redisService := auth_service.NewRedisService(redisClient)

	tokenService := auth_service.NewTokenService()

	securityService := security_service.New()

	//**********REGISTER USER*************//
	userRepository := user_repository.New(postgressClient)
	userServiceFinder := get_user_by_email_service.New(userRepository)
	createUserService := create_user_service.New(userRepository, userServiceFinder)
	createUserCommandBus := create_user_command.New(createUserService, securityService)
	create_user_controller.New(engine, createUserCommandBus, util)

	//**********LOGIN USER*************//
	loginUserService := login_user_service.New(userServiceFinder, securityService, redisService.Auth, tokenService)
	loginCreateUserQueryBus := login_user_query.New(loginUserService)
	login_user_controller.New(engine, loginCreateUserQueryBus, util)

}
