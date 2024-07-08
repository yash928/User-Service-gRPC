package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	Server          *Server
	InternalService *InternalService
}

type Server struct {
	Env  string
	Port string
}

type InternalService struct {
	UserServiceUrl string
}

func GetConfig() *Config {
	// Load the .env file
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file %#v", err)
	}

	// Get the environment variables

	server := Server{
		Env:  os.Getenv("ENVIRONMENT"),
		Port: os.Getenv("APP_PORT"),
	}

	internalService := InternalService{
		UserServiceUrl: os.Getenv("USER_SERVICE_URL"),
	}

	return &Config{
		Server:          &server,
		InternalService: &internalService,
	}
}

type CtxKey struct {
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
}
