package handler

import (
	"log"
	"net"
	userUsecase "user-service-grpc/internal/core/user"
	"user-service-grpc/internal/interfaces/input/grpc/user"

	"google.golang.org/grpc"
)

func GRPCHandler(lis net.Listener, userUC userUsecase.UserUsecase) {
	userServer := user.NewServer(userUC)
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userServer)
	err := grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve grpc server over %v", err)
	}
}
