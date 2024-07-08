package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DB              *DB
	Server          *Server
	InternalService *InternalService
}

type Server struct {
	Env      string
	Port     string
	GRPCPort string
}

type DB struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
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
	db := DB{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}

	server := Server{
		Env:      os.Getenv("ENVIRONMENT"),
		Port:     os.Getenv("APP_PORT"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}

	internalService := InternalService{
		UserServiceUrl: os.Getenv("USER_SERVICE_URL"),
	}

	return &Config{
		DB:              &db,
		Server:          &server,
		InternalService: &internalService,
	}
}

type CtxKey struct {
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
}
