package db

import (
	"database/sql"
	"fmt"
	"log"
	"user-service-grpc/config"
)

func Connect(cfg *config.DB) (*sql.DB, error) {

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Open the MySQL connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	return db, nil
}
