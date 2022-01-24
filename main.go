package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/config"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/service/middleware"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/product"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func getPostgresDBClient(config config.Interface) *gorm.DB {
	host := config.Database().Host()
	user := config.Database().User()
	pass := config.Database().Password()
	database := config.Database().Name()
	dsn := fmt.Sprintf("host=%v user=%v password=%v database=%v", host, user, pass, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func getRedisClient(config config.Interface) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis().Address(),
		Password: config.Redis().Password(),
	})
}

func main() {
	config := config.Initialize()
	utils := util.Initialize()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	apiV1 := router.Group("/api/v1/", middleware.AuthMiddleware())
	// User only can be added by authorized person
	postgressClient := getPostgresDBClient(config)
	redisClient := getRedisClient(config)
	user.Initialize(router, utils, postgressClient, redisClient)
	product.Initialize(apiV1, utils, postgressClient)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":" + config.Server().Port()))
}
