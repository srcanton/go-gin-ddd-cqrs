package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ServerConfigInterface server config interface
type ServerConfigInterface interface {
	Port() string
	Mode() string
}

// Server server config struct
type Server struct {
	port string
	mode string
}

// NewServerConfig create server config struct instance
func NewServerConfig() *Server {
	mode := "debug"
	if env := os.Getenv("GIN_MODE"); env != "" {
		mode = env
	}

	loadEnv(mode)

	port := os.Getenv("GIN_PORT")

	server := &Server{
		port: port,
		mode: mode,
	}
	if server.mode != "release" && server.mode != "debug" {
		panic("Unavailable gin mode")
	}
	return server
}

func loadEnv(mode string) {
	if mode == "debug" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
}

// Port get server port number
func (server *Server) Port() string {
	return server.port
}

// Mode get server mode
func (server *Server) Mode() string {
	return server.mode
}
