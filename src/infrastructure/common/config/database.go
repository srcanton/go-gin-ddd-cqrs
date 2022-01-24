package config

import (
	"os"
)

// DatabaseConfigInterface persistence config interface
type DatabaseConfigInterface interface {
	Host() string
	Port() string
	Name() string
	User() string
	Password() string
}

// Database persistence config struct
type Database struct {
	port     string
	host     string
	name     string
	user     string
	password string
}

// NewDatabaseConfig create persistence instance
func NewDatabaseConfig() *Database {
	port := os.Getenv("DATABASE_PORT")
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")

	database := &Database{
		port:     port,
		host:     host,
		name:     name,
		user:     user,
		password: password,
	}
	return database
}

// Host get persistence host
func (database *Database) Host() string {
	return database.host
}

// Port get persistence port number
func (database *Database) Port() string {
	return database.port
}

// Name get persistence name
func (database *Database) Name() string {
	return database.name
}

// User get databsae user name
func (database *Database) User() string {
	return database.user
}

// Password get persistence user password
func (database *Database) Password() string {
	return database.password
}
