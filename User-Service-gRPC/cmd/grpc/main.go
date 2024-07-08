package main

import (
	"fmt"
	"log"
	"net"
	"user-service-grpc/config"
	"user-service-grpc/internal/adapters/persistence/db"
	"user-service-grpc/internal/interfaces/input/grpc/user/handler"
	userUC "user-service-grpc/internal/usecase/user"
	"user-service-grpc/pkg/logging"
)

func main() {

	cfg := config.GetConfig()

	logs := logging.NewService(".logs/logger.log")

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

	userDB := db.NewUserDb(dbCon)

	userUC := userUC.NewUserUsecase(userDB, logs)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.GRPCPort))
	if err != nil {
		logs.GetLogger().Fatal("Failed to lister: %v", err)
		return
	}

	handler.GRPCHandler(lis, userUC)

}
