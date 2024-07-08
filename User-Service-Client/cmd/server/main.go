package main

import (
	"net/http"
	"user-service-client/config"

	"user-service-client/internal/interfaces/input/api"
	userHand "user-service-client/internal/interfaces/input/api/handler/user"
	"user-service-client/internal/interfaces/output/grpc/user"
	userUsecase "user-service-client/internal/usecase/user"
	"user-service-client/pkg/logging"
)

func main() {

	cfg := config.GetConfig()

	logs := logging.NewService(".logs/logger.log")

	userClient, err := user.Connect(cfg.InternalService.UserServiceUrl)
	if err != nil {
		logs.GetLogger().Fatal("Could not connect to service", err)
	}

	userUC := userUsecase.NewUserUsecase(userClient, logs)

	userHand := userHand.NewUserHand(userUC)

	r := api.SetUpRoutes(userHand)

	http.ListenAndServe(":"+cfg.Server.Port, r)

}
