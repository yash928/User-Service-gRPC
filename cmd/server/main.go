package main

import (
	"log"
	"user-service-grpc/config"
	"user-service-grpc/internal/adapters/persistence/db"
)

func main() {

	cfg := config.GetConfig()

	dbCon, err := db.Connect(cfg.DB)
	if err != nil {
		log.Print("Error=", err)
		return
	}

	defer dbCon.Close()

	// Check the connection
	err = dbCon.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return
	} else {
		log.Printf("Ping successful")
	}

}
