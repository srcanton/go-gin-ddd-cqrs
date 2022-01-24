package config

import (
	"github.com/srcanton/go-gin-ddd-cqrs/docs"
)

// Swagger swagger config struct
type Swagger struct{}

// NewSwagger create swagger configuration instance
func NewSwagger() *Swagger {
	docs.SwaggerInfo.Title = "Gin CQRS REST API"
	docs.SwaggerInfo.Description = "This is Example for REST api using Gin"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	return &Swagger{}
}
