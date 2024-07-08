package user

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(connString string) (UserServiceClient, error) {
	userServerConn, err := grpc.NewClient(connString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to dial", err)
		return nil, err
	}
	newCLient := NewUserServiceClient(userServerConn)
	return newCLient, nil
}
