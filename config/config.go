package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DB *DB
}

type DB struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
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

	return &Config{
		DB: &db,
	}
}

type CtxKey struct {
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
}
