<h1 align="center">GOLANG API REST</h1>

<p align="center">
  <a href="https://github.com/srcanton/go-gin-ddd-cqrs/actions"><img src="https://github.com/srcanton/go-gin-ddd-cqrs/workflows/test/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/srcanton/go-gin-ddd-cqrs"><img src="https://goreportcard.com/badge/github.com/srcanton/go-gin-ddd-cqrs" /></a>
  <a href="https://pkg.go.dev/github.com/srcanton/go-gin-ddd-cqrs"><img src="https://pkg.go.dev/badge/github.com/srcanton/go-gin-ddd-cqrs" /></a>
  <a href="https://github.com/srcanton/go-gin-ddd-cqrs/issues/1"><img src="https://img.shields.io/badge/chat-on%20issue-yellow"></a>
  <a href="https://github.com/srcanton/go-gin-ddd-cqrs/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPL%20v3.0-brightgreen.svg" /></a>
</p>

<p align="center">
  Clean api rest with Go, Gin and GORM.  
</p>
<p align="center">
  Clean Architecture with DDD, CQRS.
</p><p align="center">
  Use Redis for cache and Postgres for database.
</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/19743841/93784289-cbd84200-fc67-11ea-997b-b99af8affe17.png">
</p>
---

## Table of Contents

- [Getting Started](#getting-started)
- [Endpoints](#endpoints)
- [License](#license)
- [Author](#author)

## [Getting Started](#getting-started)
Develop: With local Go
```zsh
make localrun
make createdb
make migrateup
```

Release: With dockerized Go
```zsh
make run
make createdb
make migrateup
```

## Endpoints
- Public methods
  - `GET /register`
  - `GET /login`

- With authorization
  - `GET /api/v1/products`

## Package Structure

```zsh
.
├── LICENSE
├── README.md
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── docker-compose.local.yml
├── .env
├── .env.local
└── src                                       # Main Application
     ├── application
     │   ├── product
     │   │   └── query  
     │   │       └── get_products      
     │   │           ├── get_products_query.go
     │   │           ├── get_products_query_handler.go
     │   │           ├── get_products_query_response.go
     │   │           └── main.go
     │   └── user
     │       ├── command  
     │       │   ├── create_user      
     │       │   │   ├── create_user_command.go
     │       │   │   ├── create_user_command_handler.go
     │       │   │   └── main.go   
     │       └── query  
     │           ├── get_user_by_email      
     │           │   ├── get_user_by_email_query.go
     │           │   ├── get_user_by_email_query_handler.go
     │           │   ├── get_user_by_email_query_response.go
     │           │   └── main.go        
     │           └── login_user      
     │               ├── login_user_query.go
     │               ├── login_user_query_handler.go
     │               ├── login_user_query_response.go
     │               └── main.go               
     ├── domain
     │   ├── common
     │   │   ├── custom-error      
     │   │   │    └── invalid_argument_error.go
     │   │   ├── service      
     │   │   │   ├── auth      
     │   │   │   │   ├── auth.go
     │   │   │   │   └── token.go
     │   │   │   └── security      
     │   │   │       └── security.go
     │   │   └── valueobject      
     │   │       ├── access      
     │   │       │   └── access_detail.go
     │   │       └── token      
     │   │           └── token_detail.go
     │   ├── product
     │   │   ├── aggregate      
     │   │   │   └── product_aggregate.go
     │   │   ├── repository      
     │   │   │   └── product_repository.go
     │   │   ├── service      
     │   │   │   └── get_products      
     │   │   │       └── get_products.go
     │   │   └── valueobject      
     │   │       ├── cost      
     │   │       │   └── cost.go
     │   │       ├── id      
     │   │       │   └── id.go
     │   │       └── name      
     │   │           └── name.go
     │   └── user
     │       ├── aggregate      
     │       │   └── product_aggregate.go
     │       ├── errors      
     │       │   ├── already_exists_error.go
     │       │   ├── incorrect_user_or_password_error.go
     │       │   └── not_found_error.go
     │       ├── repository      
     │       │   └── user_repository.go
     │       ├── service      
     │       │   ├── create_user      
     │       │   │   └── create_user_service.go
     │       │   ├── get_user_by_email      
     │       │   │   └── get_user_by_email_service.go
     │       │   └── login_user      
     │       │       └── login_user_service.go
     │       └── valueobject      
     │           ├── email      
     │           │   └── email.go
     │           ├── first_name      
     │           │   └── first_name.go
     │           ├── id      
     │           │   └── id.go
     │           ├── last_name      
     │           │   └── last_name.go
     │           └── password      
     │               └── password.go
     └── infrastructure
         ├── common
         │   ├── config
         │   │   ├── database.go
         │   │   ├── main.go
         │   │   ├── redis.go
         │   │   ├── server.go
         │   │   └── swagger.go
         │   │
         │   ├── persistence
         │   │   ├── migration   
         │   │   └── model    
         │   │        ├── product_model.go                                
         │   │        └── user_model.go                                
         │   └── service
         │       ├── auth   
         │       │   ├── auth.go                                
         │       │   ├── redisdb.go            
         │       │   └── token.go            
         │       ├── middleware   
         │       │   └── middleware.go                     
         │       └── security    
         │           └── security.go           
         ├── product
         │   ├── controller   
         │   │   ├── get_products   
         │   │   │   ├── get_products_controller.go     
         │   │   │   └── main.go   
         │   ├── repository  
         │   │   └── product_repository  
         │   └── main.go                                                                                                                              
         └── user
             ├── controller   
             │   ├── create_user   
             │   │   ├── create_user_controller.go     
             │   │   └── main.go   
             │   └── login_user   
             │       ├── login_user_controller.go     
             │       └── main.go           
             ├── repository  
             │   └── user_repository  
             └── main.go           
```

## Feedbacks

[Feel free to write your thoughts](https://github.com/srcanton/go-gin-ddd-cqrs/issues/1)

## License

[GNU General Public License v3.0](https://github.com/srcanton/go-gin-ddd-cqrs/blob/main/LICENSE).

## Author

srcanton

<a href="https://github.com/srcanton"><img src="https://user-images.githubusercontent.com/19743841/97778118-4629a980-1bb8-11eb-97ed-76dcdbe50406.png" /></a>
<a href="https://twitter.com/sr_canton"><img src="https://user-images.githubusercontent.com/19743841/97777698-52f8ce00-1bb5-11eb-93c9-b06e0c48b693.png" /></a>
